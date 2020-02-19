package p633_1

import "math"

var JudgeSquareSum = judgeSquareSum

// 双指针
// Space: O(1)
// Runtime: O(sqrt(n))
func judgeSquareSum(c int) bool {
	if c < 0 {
		return false
	}
	a, b := 0, int(math.Ceil(math.Sqrt(float64(c))))
	for a <= b {
		sum := a*a + b*b
		if sum == c {
			return true
		}
		if sum > c {
			b--
		} else {
			a++
		}
	}
	return false
}
