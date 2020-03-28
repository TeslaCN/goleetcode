package p1162_1

func maxDistance(grid [][]int) int {
	step := 1
	for {
		modified := 0
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				value := grid[i][j]
				if value == step {
					modified += set(grid, i-1, j, step+1)
					modified += set(grid, i+1, j, step+1)
					modified += set(grid, i, j-1, step+1)
					modified += set(grid, i, j+1, step+1)
				}
			}
		}
		if modified == 0 {
			break
		}
		step++
	}
	if step == 1 {
		return -1
	}
	return step - 1
}

func set(grid [][]int, i, j, step int) int {
	if i >= 0 && j >= 0 && i < len(grid) && j < len(grid[i]) && grid[i][j] == 0 {
		grid[i][j] = step
		return 1
	}
	return 0
}
