package utils

import "loop-notes-api/internal/entities"

func CalculatePeriod(review int, score entities.Score) float64 {
	if review == 1 {
		return 30
	}
	return CalculatePeriod(review-1, score) * score.Value
}
