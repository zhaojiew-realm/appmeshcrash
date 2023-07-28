package utils

import (
	"math"
)

func CpuLoad(precision int) {
	for i := 0; i < precision; i++ {
		math.Sqrt(float64(i))
	}
}
