package p279_3

import "math"

var NumSquares = numSquares

func numSquares(n int) int {
	if isSquare(n) {
		return 1
	}
	for n%4 == 0 {
		n >>= 2
	}
	if n%8 == 7 {
		return 4
	}
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if isSquare(n - i*i) {
			return 2
		}
	}
	return 3
}

func isSquare(n int) bool {
	sqrt := math.Sqrt(float64(n))
	return sqrt == float64(int(sqrt))
}
