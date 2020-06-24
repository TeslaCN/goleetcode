package p136_1

// Runtime: O(n)
// Space: O(1)
func SingleNumber(nums []int) int {
	if nums == nil || len(nums) < 1 {
		return 0
	}
	a := 0
	for i := 0; i < len(nums); i++ {
		a ^= nums[i]
	}
	return a
}
