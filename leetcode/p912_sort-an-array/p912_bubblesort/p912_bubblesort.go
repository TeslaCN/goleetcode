package p912_bubblesort

// Deprecated: TLE
var SortArray = sortArray

// Deprecated: TLE
func sortArray(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		changed := false
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
				changed = true
			}
		}
		if !changed {
			break
		}
	}
	return nums
}
