package p240_1

func searchMatrix(matrix [][]int, target int) bool {
	x, y := len(matrix)-1, 0
	for x >= 0 && y < len(matrix[x]) {
		switch {
		case matrix[x][y] == target:
			return true
		case target < matrix[x][y]:
			x--
		case target > matrix[x][y]:
			y++
		}
	}
	return false
}
