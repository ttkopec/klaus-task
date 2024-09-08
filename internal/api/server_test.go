package api

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/ttkopec/klaus-task/internal/db/mocks"
	"github.com/ttkopec/klaus-task/internal/models"
	pb "github.com/ttkopec/klaus-task/pkg/protos/gen"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestServer(t *testing.T) {
	t.Run("GetAggregatedCategoryScores TODO", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			req := &pb.AggregatedCategoryScoreRequest{
				Period: &pb.TimeRange{
					Start: &timestamppb.Timestamp{
						Seconds: 1,
					},
					End: &timestamppb.Timestamp{
						Seconds: 2,
					},
				},
			}
			testCtx := context.Background()
			testLogger := slog.New(slog.NewTextHandler(io.Discard, nil))
			mockDb := mocks.NewMockService(t)
			mockDb.EXPECT().GetAggregatedCategoryScores(testCtx, time.Unix(1, 0).UTC(), time.Unix(2, 0).UTC()).Return(
				[]models.CategoryScorePerPeriod{
					{
						Period:   "2024-12-12",
						Score:    12.3,
						Category: "Grammar",
					},
				},
				nil,
			)

			srv := NewServer(mockDb, testLogger)

			res, err := srv.GetAggregatedCategoryScores(testCtx, req)

			require.NoError(t, err)
			require.NotNil(t, res)
			require.NotNil(t, res.Scores)
			require.Len(t, res.Scores, 1)
			require.Equal(t, res.Scores[0].Category, "Grammar")
			require.NotNil(t, res.Scores[0].ScoresPerPeriod)
			require.Len(t, res.Scores[0].ScoresPerPeriod, 1)
			require.Equal(t, res.Scores[0].ScoresPerPeriod[0].Period, "2024-12-12")
			require.Equal(t, res.Scores[0].ScoresPerPeriod[0].Score, 12.3)
		})
		t.Run("error", func(t *testing.T) {
			req := &pb.AggregatedCategoryScoreRequest{
				Period: &pb.TimeRange{
					Start: &timestamppb.Timestamp{
						Seconds: 1,
					},
					End: &timestamppb.Timestamp{
						Seconds: 2,
					},
				},
			}
			testCtx := context.Background()
			testLogger := slog.New(slog.NewTextHandler(io.Discard, nil))
			mockDb := mocks.NewMockService(t)
			mockDb.EXPECT().GetAggregatedCategoryScores(testCtx, time.Unix(1, 0).UTC(), time.Unix(2, 0).UTC()).Return(
				nil, errors.New("blah blah blah"),
			)

			srv := NewServer(mockDb, testLogger)

			res, err := srv.GetAggregatedCategoryScores(testCtx, req)

			require.EqualError(t, err, "failed to GetAggregatedCategoryScores: blah blah blah")
			require.Nil(t, res)
		})
	})

	t.Run("GetScoresByTicket", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			req := &pb.ScoresByTicketRequest{
				Period: &pb.TimeRange{
					Start: &timestamppb.Timestamp{
						Seconds: 1,
					},
					End: &timestamppb.Timestamp{
						Seconds: 2,
					},
				},
			}
			testCtx := context.Background()
			testLogger := slog.New(slog.NewTextHandler(io.Discard, nil))
			mockDb := mocks.NewMockService(t)
			mockDb.EXPECT().GetScoresByTicket(testCtx, time.Unix(1, 0).UTC(), time.Unix(2, 0).UTC()).Return(
				[]models.ScoreByTicket{
					{
						TicketID: 123,
						Score:    12.3,
						Category: "Grammar",
					},
					{
						TicketID: 123,
						Score:    45.6,
						Category: "GDPR",
					},
				},
				nil,
			)

			srv := NewServer(mockDb, testLogger)

			res, err := srv.GetScoresByTicket(testCtx, req)

			require.NoError(t, err)
			require.NotNil(t, res)
			require.NotNil(t, res.TicketScores)
			require.Len(t, res.TicketScores, 1)
			require.Equal(t, res.TicketScores[0].TicketId, "123")
			require.Equal(t, res.TicketScores[0].CategoryScores, map[string]float64{"Grammar": 12.3, "GDPR": 45.6})
		})
		t.Run("error", func(t *testing.T) {
			req := &pb.ScoresByTicketRequest{
				Period: &pb.TimeRange{
					Start: &timestamppb.Timestamp{
						Seconds: 1,
					},
					End: &timestamppb.Timestamp{
						Seconds: 2,
					},
				},
			}
			testCtx := context.Background()
			testLogger := slog.New(slog.NewTextHandler(io.Discard, nil))
			mockDb := mocks.NewMockService(t)
			mockDb.EXPECT().GetScoresByTicket(testCtx, time.Unix(1, 0).UTC(), time.Unix(2, 0).UTC()).Return(
				nil, errors.New("error when trying to query for GetScoresByTicket"),
			)
			srv := NewServer(mockDb, testLogger)

			res, err := srv.GetScoresByTicket(testCtx, req)

			require.EqualError(t, err, "failed to GetScoresByTicket: error when trying to query for GetScoresByTicket")
			require.Nil(t, res)
		})
	})

	t.Run("GetOverallQualityScore", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			req := &pb.OverallQualityScoreRequest{
				Period: &pb.TimeRange{
					Start: &timestamppb.Timestamp{
						Seconds: 3,
					},
					End: &timestamppb.Timestamp{
						Seconds: 4,
					},
				},
			}
			testCtx := context.Background()
			testLogger := slog.New(slog.NewTextHandler(io.Discard, nil))
			mockDb := mocks.NewMockService(t)
			mockDb.EXPECT().GetOverallQualityScore(testCtx, time.Unix(3, 0).UTC(), time.Unix(4, 0).UTC()).Return(
				&models.OverallQualityScore{
					Score: 90.1,
				},
				nil,
			)
			srv := NewServer(mockDb, testLogger)

			res, err := srv.GetOverallQualityScore(testCtx, req)

			require.NoError(t, err)
			require.NotNil(t, res)
			require.Equal(t, 90.1, res.OverallScore)
		})
		t.Run("error", func(t *testing.T) {
			req := &pb.OverallQualityScoreRequest{
				Period: &pb.TimeRange{
					Start: &timestamppb.Timestamp{
						Seconds: 3,
					},
					End: &timestamppb.Timestamp{
						Seconds: 4,
					},
				},
			}
			testCtx := context.Background()
			testLogger := slog.New(slog.NewTextHandler(io.Discard, nil))
			mockDb := mocks.NewMockService(t)
			mockDb.EXPECT().GetOverallQualityScore(testCtx, time.Unix(3, 0).UTC(), time.Unix(4, 0).UTC()).Return(
				nil, errors.New("internal err"),
			)
			srv := NewServer(mockDb, testLogger)

			res, err := srv.GetOverallQualityScore(testCtx, req)

			require.EqualError(t, err, "failed to GetOverallQualityScore: internal err")
			require.Nil(t, res)
		})
	})

	t.Run("GetPeriodOverPeriodScoreChange", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			req := &pb.PeriodOverPeriodScoreChangeRequest{
				CurrentPeriod: &pb.TimeRange{
					Start: &timestamppb.Timestamp{
						Seconds: 10,
					},
					End: &timestamppb.Timestamp{
						Seconds: 20,
					},
				},
				PreviousPeriod: &pb.TimeRange{
					Start: &timestamppb.Timestamp{
						Seconds: 30,
					},
					End: &timestamppb.Timestamp{
						Seconds: 40,
					},
				},
			}
			testCtx := context.Background()
			testLogger := slog.New(slog.NewTextHandler(io.Discard, nil))
			mockDb := mocks.NewMockService(t)
			mockDb.EXPECT().GetPeriodOverPeriodScoreChange(testCtx, time.Unix(10, 0).UTC(), time.Unix(20, 0).UTC(), time.Unix(30, 0).UTC(), time.Unix(40, 0).UTC()).Return(
				&models.ScoreChange{
					Change: -32.1,
				},
				nil,
			)
			srv := NewServer(mockDb, testLogger)

			res, err := srv.GetPeriodOverPeriodScoreChange(testCtx, req)

			require.NoError(t, err)
			require.NotNil(t, res)
			require.Equal(t, -32.1, res.PercentageChange)
		})
		t.Run("error", func(t *testing.T) {
			req := &pb.PeriodOverPeriodScoreChangeRequest{
				CurrentPeriod: &pb.TimeRange{
					Start: &timestamppb.Timestamp{
						Seconds: 11,
					},
					End: &timestamppb.Timestamp{
						Seconds: 22,
					},
				},
				PreviousPeriod: &pb.TimeRange{
					Start: &timestamppb.Timestamp{
						Seconds: 33,
					},
					End: &timestamppb.Timestamp{
						Seconds: 44,
					},
				},
			}
			testCtx := context.Background()
			testLogger := slog.New(slog.NewTextHandler(io.Discard, nil))
			mockDb := mocks.NewMockService(t)
			mockDb.EXPECT().GetPeriodOverPeriodScoreChange(testCtx, time.Unix(11, 0).UTC(), time.Unix(22, 0).UTC(), time.Unix(33, 0).UTC(), time.Unix(44, 0).UTC()).Return(
				nil, errors.New("random error"),
			)
			srv := NewServer(mockDb, testLogger)

			res, err := srv.GetPeriodOverPeriodScoreChange(testCtx, req)

			require.EqualError(t, err, "failed to GetPeriodOverPeriodScoreChange: random error")
			require.Nil(t, res)
		})
	})
}
