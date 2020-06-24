package p494_2

var FindTargetSumWays = findTargetSumWays

func findTargetSumWays(nums []int, S int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	cache := make([]map[int]int, n+1)
	cache[0] = map[int]int{0: 1}
	for i := 1; i <= n; i++ {
		cache[i] = map[int]int{}
		for sum, times := range cache[i-1] {
			cache[i][sum+nums[i-1]] += times
			cache[i][sum-nums[i-1]] += times
		}
	}
	return cache[n][S]
}

/*
执行用时 : 64 ms , 在所有 Go 提交中击败了 57.30% 的用户
内存消耗 : 6.5 MB , 在所有 Go 提交中击败了 50.00% 的用户
*/
