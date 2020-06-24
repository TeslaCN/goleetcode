package p581_1

var FindUnsortedSubarray = findUnsortedSubarray

// Runtime: O(n^2)
// Space: O(1)
func findUnsortedSubarray(nums []int) int {
	if nums == nil || len(nums) <= 1 {
		return 0
	}
	start, end := len(nums), -1
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				start = min(i, start)
				end = max(j, end)
			}
		}
	}
	if start > end {
		return 0
	}
	return end - start + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
