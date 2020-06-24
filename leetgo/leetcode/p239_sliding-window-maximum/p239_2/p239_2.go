package p239_2

var MaxSlidingWindow = maxSlidingWindow

// Runtime: O(n)
// Space: O(1)
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 {
		return []int{}
	}
	if k == 1 {
		return nums
	}
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		blockStart := i / k * k
		if i == blockStart {
			left[i] = nums[i]
		} else {
			left[i] = max(nums[i], left[i-1])
		}
	}
	right[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		blockEnd := i/k*k + k - 1
		if i == blockEnd {
			right[i] = nums[i]
		} else {
			right[i] = max(nums[i], right[i+1])
		}
	}
	for i := 0; i < n-k+1; i++ {
		nums[i] = max(left[i+k-1], right[i])
	}
	return nums[:n-k+1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
执行用时 : 40 ms , 在所有 Go 提交中击败了 28.66% 的用户
内存消耗 : 7 MB , 在所有 Go 提交中击败了 25.00% 的用户
*/
