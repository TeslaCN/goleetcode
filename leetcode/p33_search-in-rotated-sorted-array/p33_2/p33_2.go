package p33_2

var Search = search

// Space: O(log(n))
// Runtime: O(log(n))
func search(nums []int, target int) int {
	stack := []int{0, len(nums)}
	for len(stack) > 0 {
		from, end := stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		if from >= end {
			continue
		}
		mid := (from + end - 1) >> 1
		if nums[mid] == target {
			return mid
		}
		stack = append(stack, from, mid, mid+1, end)
	}
	return -1
}

/*
执行用时 : 4 ms , 在所有 Go 提交中击败了 73.60% 的用户
内存消耗 : 2.6 MB , 在所有 Go 提交中击败了 71.43% 的用户
*/
