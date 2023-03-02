package libs

import (
	"math"
)

func CalculateNewRating(count int64, currentRating, newRating float64) float64 {

	newRating = (currentRating*float64(count-1) + newRating) / float64(count+1)
	roundRating := math.Round(newRating*10) / 10

	return roundRating
}
