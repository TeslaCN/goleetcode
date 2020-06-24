package p69_1

import "math"

var MySqrt = mySqrt

func mySqrt(x int) int {
	return int(math.Sqrt(float64(x)))
}
