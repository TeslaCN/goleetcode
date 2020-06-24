package p53_1

// Deprecated
// Runtime: O(n^3)
// Space: O(1)
func MaxSubArray(nums []int) int {
	if nums == nil {
		return 0
	}
	if len(nums) < 1 {
		return 0
	}
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			sum := 0
			for cur := i; cur <= j; cur++ {
				sum += nums[cur]
			}
			if sum > max {
				max = sum
			}
		}
	}
	return max
}
