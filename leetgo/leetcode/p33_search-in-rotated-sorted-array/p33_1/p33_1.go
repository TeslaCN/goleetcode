package p33_1

var Search = search

// Space: O(log(n))
// Runtime: O(log(n))
func search(nums []int, target int) int {
	return bSearch(nums, 0, len(nums), target)
}

func bSearch(nums []int, from, end, target int) int {
	if from > end-1 {
		return -1
	}
	mid := (from + end - 1) >> 1
	if nums[mid] == target {
		return mid
	}
	index := bSearch(nums, from, mid, target)
	if index < 0 {
		index = bSearch(nums, mid+1, end, target)
	}
	return index
}

/*
执行用时 : 4 ms , 在所有 Go 提交中击败了 73.60% 的用户
内存消耗 : 2.6 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
