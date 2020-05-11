package p50_1

import "math"

var MyPow = myPow

func myPow(x float64, n int) float64 {
	return math.Pow(x, float64(n))
}
