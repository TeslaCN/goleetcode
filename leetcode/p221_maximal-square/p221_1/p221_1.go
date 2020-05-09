package p221_1

// Space: O(1)
// Runtime: O(m*n)
func maximalSquare(matrix [][]byte) int {
	max := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] -= 0x30
			if matrix[i][j] > 0 {
				around(matrix, i, j)
				if value := int(matrix[i][j]); value > max {
					max = value
				}
			}
		}
	}
	return max * max
}

func around(matrix [][]byte, i, j int) {
	if i < 1 || j < 1 || matrix[i-1][j] == 0 || matrix[i][j-1] == 0 || matrix[i-1][j-1] == 0 {
		return
	}
	min := matrix[i-1][j-1]
	if left := matrix[i][j-1]; left < min {
		min = left
	}
	if up := matrix[i-1][j]; up < min {
		min = up
	}
	matrix[i][j] = min + 1
}

/*
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 3.4 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
