package i17_16

func massage(nums []int) int {
	result := make([]int, len(nums)+1)
	for i := 0; i <= len(nums); i++ {
		f(nums[:i], result)
	}
	return result[len(nums)]
}

func f(nums, result []int) {
	switch len(nums) {
	case 0:
		result[0] = 0
	case 1:
		result[1] = nums[0]
	case 2:
		result[2] = max(nums[0], nums[1])
	default:
		result[len(nums)] = max(result[len(nums)-1], result[len(nums)-2]+nums[len(nums)-1])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
