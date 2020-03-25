package p892_1

var SurfaceArea = surfaceArea

func surfaceArea(grid [][]int) int {
	if len(grid) <= 0 {
		return 0
	}
	area := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			count := grid[i][j]
			if count == 0 {
				continue
			}
			area += count<<2 + 2
			if neighbor := get(grid, i-1, j); neighbor > 0 {
				area -= min(count, neighbor) << 1
			}
			if neighbor := get(grid, i, j-1); neighbor > 0 {
				area -= min(count, neighbor) << 1
			}
		}
	}
	return area
}

func get(grid [][]int, i, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) {
		return 0
	}
	return grid[i][j]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
执行用时 : 8 ms , 在所有 Go 提交中击败了 90.00% 的用户
内存消耗 : 3.7 MB , 在所有 Go 提交中击败了 40.00% 的用户
*/
