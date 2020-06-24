package p892_2

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
			if i > 0 {
				if count < grid[i-1][j] {
					area -= count << 1
				} else {
					area -= grid[i-1][j] << 1
				}
			}
			if j > 0 {
				if count < grid[i][j-1] {
					area -= count << 1
				} else {
					area -= grid[i][j-1] << 1
				}
			}
		}
	}
	return area
}
