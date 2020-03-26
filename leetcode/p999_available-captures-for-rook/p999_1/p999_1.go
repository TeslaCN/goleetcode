package p999_1

func numRookCaptures(board [][]byte) int {
	x, y := -1, -1
	for i, bytes := range board {
		for j, b := range bytes {
			if b == 'R' {
				x, y = i, j
				break
			}
		}
	}
	count := 0
	for i := x - 1; i >= 0; i-- {
		if board[i][y] == 'p' {
			count++
			break
		}
		if board[i][y] == 'B' {
			break
		}
	}
	for i := x + 1; i < len(board); i++ {
		if board[i][y] == 'p' {
			count++
			break
		}
		if board[i][y] == 'B' {
			break
		}
	}
	for i := y - 1; i >= 0; i-- {
		if board[x][i] == 'p' {
			count++
			break
		}
		if board[x][i] == 'B' {
			break
		}
	}
	for i := y + 1; i < len(board[x]); i++ {
		if board[x][i] == 'p' {
			count++
			break
		}
		if board[x][i] == 'B' {
			break
		}
	}
	return count
}
