package utils

import (
	"math"
)

func Log(x float64, base float64) float64 {
	return math.Log(x) / math.Log(base)
}

func Round(x float64) int {
	return int(math.Round(x))
}