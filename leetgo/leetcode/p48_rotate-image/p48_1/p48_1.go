package p48_1

var Rotate = rotate

/*
...        [row][col]         ...
[n-1-col][row] ... [col][n-1-row]
...    [n-1-row][n-1-col]     ...
*/
func rotate(matrix [][]int) {
	n := len(matrix)
	tmp := 0
	for row := 0; row < n/2; row++ {
		for col := row; col < n-row-1; col++ {
			tmp = matrix[row][col]
			matrix[row][col] = matrix[n-1-col][row]
			matrix[n-1-col][row] = matrix[n-1-row][n-1-col]
			matrix[n-1-row][n-1-col] = matrix[col][n-1-row]
			matrix[col][n-1-row] = tmp
		}
	}
}
