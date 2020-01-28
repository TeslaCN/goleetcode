package p53_2

// Runtime: O(n)
// Space: O(1)
func MaxSubArray(nums []int) int {
	if nums == nil {
		return 0
	}
	if len(nums) < 1 {
		return 0
	}
	max := nums[0]

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > max {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}
