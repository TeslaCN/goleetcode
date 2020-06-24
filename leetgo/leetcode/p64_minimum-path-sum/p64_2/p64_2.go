package p64_2

var MinPathSum = minPathSum

// Runtime: O(mn)
// Space: O(1)
func minPathSum(grid [][]int) int {
	n := len(grid)
	if n <= 0 {
		return 0
	}
	m := len(grid[0])
	if m <= 0 {
		return 0
	}
	for i := 1; i < n; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for j := 1; j < m; j++ {
		grid[0][j] += grid[0][j-1]
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[n-1][m-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
执行用时 : 8 ms , 在所有 Go 提交中击败了 92.63% 的用户
内存消耗 : 3.9 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
