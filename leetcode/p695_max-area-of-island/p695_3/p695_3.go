package p695_3

var MaxAreaOfIsland = maxAreaOfIsland

func maxAreaOfIsland(grid [][]int) int {
	max := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			area := dfs(i, j, grid)
			if area > max {
				max = area
			}
		}
	}
	return max
}

func dfs(x, y int, grid [][]int) int {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[x]) || grid[x][y] == 0 {
		return 0
	}
	grid[x][y] = 0
	return 1 + dfs(x+1, y, grid) + dfs(x-1, y, grid) + dfs(x, y+1, grid) + dfs(x, y-1, grid)
}

/*
dfs函数定义在 maxAreaOfIsland 内部
执行用时 : 20 ms , 在所有 Go 提交中击败了 26.26% 的用户
内存消耗 : 5 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/

/*
dfs函数定义在外部

执行用时 : 16 ms , 在所有 Go 提交中击败了 73.23% 的用户
内存消耗 : 5 MB , 在所有 Go 提交中击败了 91.67% 的用户

执行用时 : 12 ms , 在所有 Go 提交中击败了 95.45% 的用户
内存消耗 : 5 MB , 在所有 Go 提交中击败了 91.67% 的用户
*/
