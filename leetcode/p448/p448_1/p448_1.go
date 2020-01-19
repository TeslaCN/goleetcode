package p448_1

// Runtime: O(n)
// Space: O(n)
// 使用 hash 标记已有的数字，需要额外空间
func FindDisappearedNumbers(nums []int) []int {
	if nums == nil || len(nums) < 1 {
		return nums
	}
	size := len(nums)
	ints := make([]int, size)
	for i := 0; i < size; i++ {
		ints[i] = i + 1
	}
	count := 0
	for i := 0; i < size; i++ {
		j := nums[i] - 1
		if ints[j] != -1 {
			count++
			ints[j] = -1
		}
	}
	result := make([]int, size-count)
	j := 0
	for i := 0; i < size; i++ {
		if ints[i] != -1 {
			result[j] = ints[i]
			j++
		}
	}
	return result
}
