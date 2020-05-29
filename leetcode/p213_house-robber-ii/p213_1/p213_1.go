package p213_1

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return max(f(nums[:len(nums)-1]), f(nums[1:]))
}

func f(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	cache := make([]int, len(nums))
	cache[0], cache[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		cache[i] = max(cache[i-2]+nums[i], cache[i-1])
	}
	return cache[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 2.1 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
