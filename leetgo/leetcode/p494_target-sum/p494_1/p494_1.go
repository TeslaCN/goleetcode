package p494_1

var FindTargetSumWays = findTargetSumWays

func findTargetSumWays(nums []int, S int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		if nums[0] == S && nums[0] == -S {
			return 2
		}
		if nums[0] == S || nums[0] == -S {
			return 1
		}
		return 0
	}
	return findTargetSumWays(nums[:n-1], S-nums[n-1]) + findTargetSumWays(nums[:n-1], S+nums[n-1])
}

/*
执行用时 : 380 ms , 在所有 Go 提交中击败了 43.06% 的用户
内存消耗 : 2.1 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
