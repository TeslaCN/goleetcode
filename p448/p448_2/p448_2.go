package p448_2

// Runtime: O(n)
// Space: O(1)
func FindDisappearedNumbers(nums []int) []int {
	if nums == nil || len(nums) < 1 {
		return nums
	}
	size := len(nums)
	for i := 0; i < size; i++ {
		if nums[i] > 0 {
			setZeroRecursive(nums, nums[i]-1)
		}
	}
	count := 0
	for i := 0; i < size; i++ {
		if nums[i] > 0 {
			count++
		}
	}
	results := make([]int, count)
	j := 0
	for i := 0; i < size; i++ {
		if nums[i] > 0 {
			results[j] = i + 1
			j++
		}
	}
	return results
}

func setZeroRecursive(nums []int, i int) {
	next := nums[i]
	if next > 0 {
		nums[i] = 0
		setZeroRecursive(nums, next-1)
	}
}
