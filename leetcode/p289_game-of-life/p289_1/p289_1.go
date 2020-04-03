package p289_1

func gameOfLife(board [][]int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			around := 0
			for x := i - 1; x <= i+1; x++ {
				for y := j - 1; y <= j+1; y++ {
					if x < 0 || y < 0 || x >= len(board) || y >= len(board[x]) || x == i && y == j {
						continue
					}
					value := board[x][y]
					if value >= 8 || value > 0 && (x > i || x == i && y == j+1) {
						around++
					}
				}
			}
			switch board[i][j] {
			case 0:
				board[i][j] = -around
			case 1:
				board[i][j] = 8 + around
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			switch board[i][j] {
			case -3:
				fallthrough
			case 10:
				fallthrough
			case 11:
				board[i][j] = 1
			default:
				board[i][j] = 0
			}
		}
	}
}

/*
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 2 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
