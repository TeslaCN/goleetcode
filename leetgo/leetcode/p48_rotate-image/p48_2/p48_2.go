package p48_2

var Rotate = rotate

func rotate(matrix [][]int) {
	n := len(matrix)
	for row := 0; row < n/2; row++ {
		for col := row; col < n-row-1; col++ {
			matrix[row][col], matrix[n-1-col][row], matrix[n-1-row][n-1-col], matrix[col][n-1-row] =
				matrix[n-1-col][row], matrix[n-1-row][n-1-col], matrix[col][n-1-row], matrix[row][col]
		}
	}
}
