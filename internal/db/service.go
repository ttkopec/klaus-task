package db

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ttkopec/klaus-task/internal/config"
	"github.com/ttkopec/klaus-task/internal/models"
)

const UTCWithoutTimezoneFormat = "2006-01-02T15:04:05"

type Service interface {
	GetAggregatedCategoryScores(ctx context.Context, start, end time.Time) ([]models.CategoryScorePerPeriod, error)
	GetScoresByTicket(ctx context.Context, start, end time.Time) ([]models.ScoreByTicket, error)
	GetOverallQualityScore(ctx context.Context, start, end time.Time) (*models.OverallQualityScore, error)
	GetPeriodOverPeriodScoreChange(ctx context.Context, start1, end1, start2, end2 time.Time) (*models.ScoreChange, error)
}

type SqlLiteService struct {
	db     *sql.DB
	logger *slog.Logger
}

func (s *SqlLiteService) GetAggregatedCategoryScores(ctx context.Context, start, end time.Time) ([]models.CategoryScorePerPeriod, error) {
	startTime, endTime := s.formatTime(start, end)

	s.logger.Debug("querying aggregatedCategoryScoresQuery", slog.String("start", startTime), slog.String("end", endTime))

	rows, err := s.db.QueryContext(ctx, aggregatedCategoryScoresQuery, startTime, endTime, endTime, startTime)
	if err != nil {
		return nil, fmt.Errorf("failed to query aggregatedCategoryScoresQuery: %w", err)
	}
	defer rows.Close()

	var scores []models.CategoryScorePerPeriod
	for rows.Next() {
		var result models.CategoryScorePerPeriod
		err := rows.Scan(&result.Category, &result.Period, &result.Score)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		scores = append(scores, result)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error with rows: %w", err)
	}

	return scores, nil
}

func (s *SqlLiteService) GetScoresByTicket(ctx context.Context, start, end time.Time) ([]models.ScoreByTicket, error) {
	startTime, endTime := s.formatTime(start, end)

	s.logger.Debug("querying scoresByTicketQuery", slog.String("start", startTime), slog.String("end", endTime))

	rows, err := s.db.QueryContext(ctx, scoresByTicketQuery, startTime, endTime)
	if err != nil {
		return nil, fmt.Errorf("failed to query scoresByTicketQuery: %w", err)
	}
	defer rows.Close()

	var scoresByTicket []models.ScoreByTicket
	for rows.Next() {
		var result models.ScoreByTicket
		err := rows.Scan(&result.TicketID, &result.Score, &result.Category)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		scoresByTicket = append(scoresByTicket, result)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error with rows: %w", err)
	}

	return scoresByTicket, nil
}

func (s *SqlLiteService) GetOverallQualityScore(ctx context.Context, start, end time.Time) (*models.OverallQualityScore, error) {
	startTime, endTime := s.formatTime(start, end)

	s.logger.Debug("querying overallQualityScoreQuery", slog.String("start", startTime), slog.String("end", endTime))

	score := &models.OverallQualityScore{}
	err := s.db.QueryRowContext(ctx, overallQualityScoreQuery, startTime, endTime).Scan(&score.Score)
	if err != nil {
		if err == sql.ErrNoRows {
			s.logger.Debug("no rows were returned for overallQualityScoreQuery")
		}
		return nil, fmt.Errorf("failed to query overallQualityScoreQuery: %w", err)
	}

	return score, nil
}

func (s *SqlLiteService) GetPeriodOverPeriodScoreChange(ctx context.Context, start1, end1, start2, end2 time.Time) (*models.ScoreChange, error) {
	startTime1, endTime1 := s.formatTime(start1, end1)
	startTime2, endTime2 := s.formatTime(start2, end2)

	s.logger.Debug(
		"querying periodOverPeriodScoreChange",
		slog.String("start1", startTime1), slog.String("end1", endTime1),
		slog.String("start2", startTime2), slog.String("end2", endTime2),
	)

	scoreChange := &models.ScoreChange{}
	err := s.db.QueryRowContext(ctx, periodOverPeriodScoreChangeQuery, startTime2, endTime2, startTime1, endTime1).Scan(&scoreChange.Change)
	if err != nil {
		if err == sql.ErrNoRows {
			s.logger.Debug("no rows were returned for periodOverPeriodScoreChange")
		}
		return nil, fmt.Errorf("failed to query periodOverPeriodScoreChange: %w", err)
	}

	return scoreChange, nil
}

func (s *SqlLiteService) Close() {
	if s.db != nil {
		err := s.db.Close()
		if err != nil {
			s.logger.Error("failed to close database", slog.String("error", err.Error()))
		}
	}
}

func (s *SqlLiteService) formatTime(start, end time.Time) (string, string) {
	return start.Local().Format(UTCWithoutTimezoneFormat), end.Local().Format(UTCWithoutTimezoneFormat)
}

func (s *SqlLiteService) init(cfg config.Config) error {
	db, err := sql.Open("sqlite3", cfg.DbName)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	s.db = db

	s.logger.Info("connected to database")

	return nil
}

func NewService(cfg config.Config, logger *slog.Logger) (*SqlLiteService, error) {
	new := &SqlLiteService{
		logger: logger,
	}
	err := new.init(cfg)

	return new, err
}
