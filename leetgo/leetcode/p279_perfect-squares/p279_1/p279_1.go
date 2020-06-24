package p279_1

import "math"

var NumSquares = numSquares

// Deprecated: BF
func numSquares(n int) int {
	sqrt := math.Sqrt(float64(n))
	if sqrt == float64(int(sqrt)) {
		return 1
	}
	sqrtInt := int(sqrt)
	min := numSquares(n - sqrtInt*sqrtInt)
	for sub := sqrtInt - 1; sub >= 1; sub-- {
		res := numSquares(n - sub*sub)
		if res < min {
			min = res
		}
	}
	return min + 1
}
