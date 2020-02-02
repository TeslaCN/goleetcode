package p581_2

var FindUnsortedSubarray = findUnsortedSubarray

// Runtime: O(n)
// Space: O(1)
func findUnsortedSubarray(nums []int) int {
	if nums == nil || len(nums) <= 1 {
		return 0
	}
	length := len(nums)
	start, end := -1, -1
	maxVal, minVal := nums[0], nums[length-1]
	for i := 0; i < length; i++ {
		if maxVal > nums[i] {
			end = i
		} else {
			maxVal = nums[i]
		}
		if minVal < nums[length-i-1] {
			start = length - i - 1
		} else {
			minVal = nums[length-i-1]
		}
	}
	if end > start {
		return end - start + 1
	}
	return 0
}
