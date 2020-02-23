package p69_2

var MySqrt = mySqrt

// binary search
// Runtime: O(log(n))
// Space: O(1)
func mySqrt(x int) int {
	if x <= 0 {
		return 0
	}
	left, right := 1, x/2
	for left != right {
		mid := (left + right + 1) / 2
		square := mid * mid
		if square == x {
			return mid
		}
		if square > x {
			if mid == right {
				right--
			} else {
				right = mid
			}
		} else {
			if mid == left {
				left++
			} else {
				left = mid
			}
		}
	}
	return left
}
