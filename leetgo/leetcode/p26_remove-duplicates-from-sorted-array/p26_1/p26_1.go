package p26_1

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	target, source := 1, 1
	for source < len(nums) {
		if nums[source] == nums[target-1] {
			source++
		} else {
			nums[target] = nums[source]
			source++
			target++
		}
	}
	return target
}
