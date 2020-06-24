package p200_1

func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if setZero(grid, i, j) > 0 {
				count++
			}
		}
	}
	return count
}

func setZero(grid [][]byte, i, j int) int {
	if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[i]) && grid[i][j] == '1' {
		grid[i][j] = '0'
		return 1 + setZero(grid, i-1, j) + setZero(grid, i+1, j) + setZero(grid, i, j-1) + setZero(grid, i, j+1)
	}
	return 0
}

/*
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 2.9 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
