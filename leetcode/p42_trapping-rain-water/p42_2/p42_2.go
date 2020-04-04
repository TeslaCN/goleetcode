package p42_2

var Trap = trap

func trap(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}
	fromLeft, fromRight := make([]int, n), make([]int, n)
	maxLeft, maxRight := height[0], height[n-1]
	fromLeft[0], fromRight[n-1] = maxLeft, maxRight
	sum := 0
	for i := 1; i < n; i++ {
		if height[i] > maxLeft {
			maxLeft = height[i]
		}
		fromLeft[i] = maxLeft
		if height[n-1-i] > maxRight {
			maxRight = height[n-1-i]
		}
		fromRight[n-1-i] = maxRight
	}
	for i := 0; i < n; i++ {
		min := fromLeft[i]
		if fromRight[i] < min {
			min = fromRight[i]
		}
		sum += min - height[i]
	}
	return sum
}

/*
执行用时 : 4 ms , 在所有 Go 提交中击败了 86.97% 的用户
内存消耗 : 3.1 MB , 在所有 Go 提交中击败了 11.11% 的用户
*/
