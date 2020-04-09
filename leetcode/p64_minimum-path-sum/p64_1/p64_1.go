package p64_1

var MinPathSum = minPathSum

// Runtime: O(mn)
// Space: O(mn)
func minPathSum(grid [][]int) int {
	n := len(grid)
	if n <= 0 {
		return 0
	}
	m := len(grid[0])
	if m <= 0 {
		return 0
	}
	cache := make([][]int, n)
	for i := 0; i < n; i++ {
		cache[i] = make([]int, m)
	}
	cache[0][0] = grid[0][0]
	for i := 1; i < n; i++ {
		cache[i][0] = cache[i-1][0] + grid[i][0]
	}
	for j := 1; j < m; j++ {
		cache[0][j] = cache[0][j-1] + grid[0][j]
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			cache[i][j] = min(cache[i-1][j], cache[i][j-1]) + grid[i][j]
		}
	}
	return cache[n-1][m-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
执行用时 : 8 ms , 在所有 Go 提交中击败了 92.86% 的用户
内存消耗 : 4.4 MB , 在所有 Go 提交中击败了 14.81% 的用户

执行用时 : 12 ms , 在所有 Go 提交中击败了 23.62% 的用户
内存消耗 : 4.4 MB , 在所有 Go 提交中击败了 29.63% 的用户
*/
