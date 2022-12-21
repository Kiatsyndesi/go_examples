package Looking_for_a_benefactor

import (
	"math"
)

func NewAvg(arr []float64, navg float64) int64 {
	result := 0.0
	result = navg * (float64(len(arr)) + 1.0) - sum(arr)

	if result <= 0 {
		return -1
	}

	return int64(math.RoundToEven(result))
}

func sum(arr []float64) (sum float64) {
	for _, value := range arr {
		sum += value
	}
	return sum
}