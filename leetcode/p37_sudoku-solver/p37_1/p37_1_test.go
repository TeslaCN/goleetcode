package p37_1

import (
	p37_sudoku_solver "github.com/TeslaCN/goleetcode/leetcode/p37_sudoku-solver"
	"testing"
)

func Test_solveSudoku(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
	}{
		{"sample-0", args{[][]byte{
			{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			before := p37_sudoku_solver.FormatBoard(tt.args.board)
			t.Log(before)
			solveSudoku(tt.args.board)
			after := p37_sudoku_solver.FormatBoard(tt.args.board)
			t.Log(after)
		})
	}
}
