package p37_sudoku_solver

import "strings"

func FormatBoard(board [][]byte) string {
	builder := strings.Builder{}
	builder.WriteByte('\n')
	for i, bytes := range board {
		for j, b := range bytes {
			builder.WriteByte(b)
			builder.WriteByte(' ')
			if j%3 == 2 && j < 8 {
				builder.Write([]byte("| "))
			}
		}
		builder.WriteByte('\n')
		if i%3 == 2 && i < 8 {
			builder.Write([]byte(strings.Repeat("-", 21) + "\n"))
		}
	}
	return builder.String()
}
