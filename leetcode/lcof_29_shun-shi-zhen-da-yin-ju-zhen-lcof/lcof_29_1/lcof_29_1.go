package lcof_29_1

func spiralOrder(matrix [][]int) []int {
	row := len(matrix)
	if row == 0 {
		return []int{}
	}
	min, column := row, len(matrix[0])
	ans := make([]int, 0, row*column)
	if column < min {
		min = column
	}
	for i := 0; i < (min+1)>>1; i++ {
		for c := i; c < column-i; c++ {
			ans = append(ans, matrix[i][c])
		}
		for r := i + 1; r < row-i; r++ {
			ans = append(ans, matrix[r][column-1-i])
		}
		for c := column - 2 - i; i != row-1-i && c > i; c-- {
			ans = append(ans, matrix[row-1-i][c])
		}
		for r := row - 1 - i; i != column-1-i && r > i; r-- {
			ans = append(ans, matrix[r][i])
		}
	}
	return ans
}

/*
执行用时 : 8 ms , 在所有 Go 提交中击败了 99.75% 的用户
内存消耗 : 6 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
