package api

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"sync"

	"github.com/ttkopec/klaus-task/internal/config"
	"github.com/ttkopec/klaus-task/internal/db"
	"github.com/ttkopec/klaus-task/internal/middleware"
	"github.com/ttkopec/klaus-task/internal/models"
	pb "github.com/ttkopec/klaus-task/pkg/protos/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	db     db.Service
	logger *slog.Logger
	pb.UnimplementedScoreServiceServer
}

func (s *server) GetAggregatedCategoryScores(ctx context.Context, req *pb.AggregatedCategoryScoreRequest) (*pb.AggregatedCategoryScoresResponse, error) {
	scores, err := s.db.GetAggregatedCategoryScores(ctx, req.Period.Start.AsTime(), req.Period.End.AsTime())
	if err != nil {
		return nil, fmt.Errorf("failed to GetAggregatedCategoryScores: %w", err)
	}

	result := &pb.AggregatedCategoryScoresResponse{
		Scores: models.GetScoresPerPeriod(scores),
	}

	return result, nil
}

func (s *server) GetScoresByTicket(ctx context.Context, req *pb.ScoresByTicketRequest) (*pb.ScoresByTicketResponse, error) {
	scores, err := s.db.GetScoresByTicket(ctx, req.Period.Start.AsTime(), req.Period.End.AsTime())
	if err != nil {
		return nil, fmt.Errorf("failed to GetScoresByTicket: %w", err)
	}

	result := &pb.ScoresByTicketResponse{
		TicketScores: models.GetTicketScores(scores),
	}

	return result, nil
}

func (s *server) GetOverallQualityScore(ctx context.Context, req *pb.OverallQualityScoreRequest) (*pb.OverallQualityScoreResponse, error) {
	score, err := s.db.GetOverallQualityScore(ctx, req.Period.Start.AsTime(), req.Period.End.AsTime())
	if err != nil {
		return nil, fmt.Errorf("failed to GetOverallQualityScore: %w", err)
	}

	result := &pb.OverallQualityScoreResponse{
		OverallScore: score.Score,
	}

	return result, nil
}

func (s *server) GetPeriodOverPeriodScoreChange(ctx context.Context, req *pb.PeriodOverPeriodScoreChangeRequest) (*pb.PeriodOverPeriodScoreChangeResponse, error) {
	scoreChange, err := s.db.GetPeriodOverPeriodScoreChange(ctx, req.CurrentPeriod.Start.AsTime(), req.CurrentPeriod.End.AsTime(), req.PreviousPeriod.Start.AsTime(), req.PreviousPeriod.End.AsTime())
	if err != nil {
		return nil, fmt.Errorf("failed to GetPeriodOverPeriodScoreChange: %w", err)
	}
	result := &pb.PeriodOverPeriodScoreChangeResponse{
		PercentageChange: scoreChange.Change,
	}

	return result, nil
}

func NewServer(db db.Service, logger *slog.Logger) *server {
	return &server{
		db:     db,
		logger: logger,
	}
}

func Start(ctx context.Context, srv *server, cfg config.Config, logger *slog.Logger) {
	grpcSrv := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.ErrorLoggingInterceptor(logger)),
	)
	pb.RegisterScoreServiceServer(grpcSrv, srv)

	// Register reflection service on gRPC server if in debug mode
	if cfg.IsDebugMode {
		reflection.Register(grpcSrv)
	}

	go func() {
		listener, err := net.Listen("tcp", cfg.GetAddress())
		if err != nil {
			logger.Error("failed to listen", slog.String("error", err.Error()))
			return
		}

		logger.Info("started listening", slog.String("address", listener.Addr().String()))

		if err := grpcSrv.Serve(listener); err != nil {
			logger.Error("failed to serve", slog.String("error", err.Error()))
		}

	}()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		<-ctx.Done()

		grpcSrv.Stop()

		logger.Info("server stopped")
	}()

	wg.Wait()

}
