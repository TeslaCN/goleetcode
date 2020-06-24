package p62_1

var UniquePaths = uniquePaths

func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	cache := make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
		cache[i][0] = 1
	}
	for i := 0; i < n; i++ {
		cache[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			cache[i][j] = cache[i-1][j] + cache[i][j-1]
		}
	}
	return cache[m-1][n-1]
}

/*
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 2 MB , 在所有 Go 提交中击败了 73.91% 的用户
*/
