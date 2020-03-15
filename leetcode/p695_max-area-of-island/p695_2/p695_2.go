package p695_2

var MaxAreaOfIsland = maxAreaOfIsland

func maxAreaOfIsland(grid [][]int) int {
	visited := make([][]int, len(grid))
	max := 0
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]int, len(grid[i]))
	}

	var dfs func(int, int) int
	dfs = func(x, y int) int {
		if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[x]) || grid[x][y] == 0 || visited[x][y] == 1 {
			return 0
		}
		visited[x][y] = 1
		return 1 + dfs(x+1, y) + dfs(x-1, y) + dfs(x, y+1) + dfs(x, y-1)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 || visited[i][j] == 1 {
				continue
			}
			area := dfs(i, j)
			if area > max {
				max = area
			}
		}
	}
	return max
}

/*
执行用时 : 12 ms , 在所有 Go 提交中击败了 95.45% 的用户
内存消耗 : 5.6 MB , 在所有 Go 提交中击败了 12.50% 的用户
*/
