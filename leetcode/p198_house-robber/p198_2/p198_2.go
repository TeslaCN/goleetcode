package p198_2

func Rob(nums []int) int {
	if nums == nil {
		return 0
	}
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		return max(nums[0], nums[1])
	default:
		f := make([]int, len(nums))
		f[0] = nums[0]
		f[1] = max(f[0], nums[1])
		f[2] = max(nums[1], f[0]+nums[2])
		for i := 3; i < len(f); i++ {
			f[i] = max(max(f[i-2], f[i-3])+nums[i], f[i-1])
		}
		return f[len(f)-1]
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
