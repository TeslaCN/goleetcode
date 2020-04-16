package p542_1

import "math"

func updateMatrix(matrix [][]int) [][]int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 1 {
				matrix[i][j] = math.MaxInt32
			}
		}
	}
	base := 0
	for {
		changed := false
		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[i]); j++ {
				if matrix[i][j] == base {
					set(matrix, i-1, j, math.MaxInt32, base+1)
					set(matrix, i+1, j, math.MaxInt32, base+1)
					set(matrix, i, j-1, math.MaxInt32, base+1)
					set(matrix, i, j+1, math.MaxInt32, base+1)
					changed = true
				}
			}
		}
		base++
		if !changed {
			break
		}
	}
	return matrix
}

func set(matrix [][]int, i, j, expect, target int) {
	if i >= 0 && i < len(matrix) && j >= 0 && j < len(matrix[i]) && matrix[i][j] == expect {
		matrix[i][j] = target
	}
}

/*
执行用时 : 64 ms , 在所有 Go 提交中击败了 91.14% 的用户
内存消耗 : 6.6 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
