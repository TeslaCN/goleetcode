package p695_1

var MaxAreaOfIsland = maxAreaOfIsland

func maxAreaOfIsland(grid [][]int) int {
	visited := map[[2]int]struct{}{}
	max := 0
	for i, nums := range grid {
		for j, num := range nums {
			if _, exists := visited[[2]int{i, j}]; num != 0 && !exists {
				island := map[[2]int]struct{}{}
				dfs(visited, island, [2]int{i, j}, grid)
				if len(island) > max {
					max = len(island)
				}
			}
		}
	}
	return max
}

func dfs(visited, island map[[2]int]struct{}, point [2]int, grid [][]int) {
	i := point[0]
	j := point[1]
	if _, exists := visited[point]; exists || i < 0 || i > len(grid)-1 || j < 0 || j > len(grid[i])-1 || grid[i][j] == 0 {
		return
	}
	visited[point] = struct{}{}
	island[point] = struct{}{}
	dfs(visited, island, [2]int{i - 1, j}, grid)
	dfs(visited, island, [2]int{i + 1, j}, grid)
	dfs(visited, island, [2]int{i, j - 1}, grid)
	dfs(visited, island, [2]int{i, j + 1}, grid)
}

/*
执行用时 : 28 ms , 在所有 Go 提交中击败了 12.63% 的用户
内存消耗 : 6.1 MB , 在所有 Go 提交中击败了 6.25% 的用户
*/
