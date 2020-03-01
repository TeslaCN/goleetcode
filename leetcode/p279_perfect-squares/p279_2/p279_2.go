package p279_2

import "math"

var NumSquares = numSquares

var cache = make(map[int]int)

func numSquares(n int) int {
	if val, ok := cache[n]; ok {
		return val
	}
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
	cache[n] = min + 1
	return cache[n]
}
