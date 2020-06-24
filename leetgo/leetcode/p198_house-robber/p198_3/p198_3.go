package p198_3

var Rob = rob

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	nums[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		nums[i] = max(nums[i-2]+nums[i], nums[i-1])
	}
	return nums[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 2 MB , 在所有 Go 提交中击败了 33.33% 的用户
*/
