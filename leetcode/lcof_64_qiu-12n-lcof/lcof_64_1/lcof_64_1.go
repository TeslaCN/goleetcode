package lcof_64_1

import "math"

// Deprecated
var SumNums = sumNums

func sumNums(n int) int {
	return (int(math.Pow(float64(n), 2)) + n) >> 1
}
