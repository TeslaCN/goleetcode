package p494_3

var FindTargetSumWays = findTargetSumWays

func findTargetSumWays(nums []int, S int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, 2001)
	dp[nums[0]+1000] = 1
	dp[-nums[0]+1000] += 1
	for i := 1; i < n; i++ {
		next := make([]int, 2001)
		for sum := -1000; sum <= 1000; sum++ {
			if dp[sum+1000] > 0 {
				next[sum+nums[i]+1000] += dp[sum+1000]
				next[sum-nums[i]+1000] += dp[sum+1000]
			}
		}
		dp = next
	}
	if S < -1000 || S > 1000 {
		return 0
	}
	return dp[S+1000]
}

/*
执行用时 : 12 ms , 在所有 Go 提交中击败了 71.17% 的用户
内存消耗 : 6.4 MB , 在所有 Go 提交中击败了 50.00% 的用户
*/
