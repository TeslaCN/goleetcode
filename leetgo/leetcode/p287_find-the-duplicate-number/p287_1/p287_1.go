package p287_1

var FindDuplicate = findDuplicate

// Runtime: O(n^2)
// Space: O(1)
func findDuplicate(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return nums[i]
			}
		}
	}
	return 0
}
