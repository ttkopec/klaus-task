package models

import (
	"strconv"

	pb "github.com/ttkopec/klaus-task/pkg/protos/gen"
)

type CategoryScorePerPeriod struct {
	Category string
	Period   string
	Score    float64
}

func GetScoresPerPeriod(in []CategoryScorePerPeriod) []*pb.CategoryScores {
	resultMap := make(map[string]*pb.CategoryScores, 0)

	for _, score := range in {
		if val, ok := resultMap[score.Category]; !ok {
			resultMap[score.Category] = &pb.CategoryScores{
				Category: score.Category,
				ScoresPerPeriod: []*pb.ScorePerPeriod{
					{
						Period: score.Period,
						Score:  score.Score,
					},
				},
			}
		} else {
			scorePerPeriod := &pb.ScorePerPeriod{
				Period: score.Period,
				Score:  score.Score,
			}
			val.ScoresPerPeriod = append(val.ScoresPerPeriod, scorePerPeriod)
		}
	}

	resultsSlice := make([]*pb.CategoryScores, 0, 4)
	for _, v := range resultMap {
		resultsSlice = append(resultsSlice, v)
	}

	return resultsSlice
}

type ScoreByTicket struct {
	TicketID int
	Score    float64
	Category string
}

func GetTicketScores(in []ScoreByTicket) []*pb.TicketScore {
	result := make([]*pb.TicketScore, 0, len(in)/4)
	intermediaryMap := make(map[int]*pb.TicketScore, len(in)/4)

	for _, score := range in {
		if val, ok := intermediaryMap[score.TicketID]; !ok {
			categoryScores := make(map[string]float64, 4)
			categoryScores[score.Category] = score.Score

			intermediaryMap[score.TicketID] = &pb.TicketScore{
				TicketId:       strconv.Itoa(score.TicketID),
				CategoryScores: categoryScores,
			}
		} else {
			val.CategoryScores[score.Category] = score.Score
		}
	}

	for _, val := range intermediaryMap {
		result = append(result, val)
	}

	return result
}

type OverallQualityScore struct {
	Score float64
}

type ScoreChange struct {
	Change float64
}
