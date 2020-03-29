package p37_1

func solveSudoku(board [][]byte) {
	cache := [3][3][9]bool{}
	slot := make([][2]int, 0)
	for i, bytes := range board {
		for j, b := range bytes {
			if b != '.' {
				cache[i/3][j/3][b-0x31] = true
			} else {
				slot = append(slot, [2]int{i, j})
			}
		}
	}
	backTrack(board, slot, 0, cache)
}

const base byte = 0x31

func backTrack(board [][]byte, slot [][2]int, level int, cache [3][3][9]bool) bool {
	if level == len(slot) {
		return true
	} else {
		x, y := slot[level][0], slot[level][1]
		for i := byte(0); i < 9; i++ {
			if cache[x/3][y/3][i] {
				continue
			}
			board[x][y] = i + base
			if valid(board, x, y, cache) {
				cache[x/3][y/3][i] = true
				if backTrack(board, slot, level+1, cache) {
					return true
				}
				cache[x/3][y/3][i] = false
			}
			board[x][y] = '.'
		}
		return false
	}
}

func valid(board [][]byte, i, j int, cache [3][3][9]bool) bool {
	b := board[i][j]
	if cache[i/3][j/3][b-base] {
		return false
	}
	for a := 0; a < 9; a++ {
		if a != i && board[a][j] == board[i][j] {
			return false
		}
	}
	for a := 0; a < 9; a++ {
		if a != j && board[i][a] == board[i][j] {
			return false
		}
	}
	return true
}

/*
执行用时 : 4 ms , 在所有 Go 提交中击败了 71.51% 的用户
内存消耗 : 2.1 MB , 在所有 Go 提交中击败了 39.13% 的用户
*/
