package p51_1

import (
	"reflect"
	"testing"
)

func Test_solveNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{"sample-0", args{4}, [][]string{
			{
				".Q..",
				"...Q",
				"Q...",
				"..Q.",
			},
			{
				"..Q.",
				"Q...",
				"...Q",
				".Q..",
			},
		}},
		{"sample-1", args{6}, [][]string{
			{".Q....", "...Q..", ".....Q", "Q.....", "..Q...", "....Q."}, {"..Q...", ".....Q", ".Q....", "....Q.", "Q.....", "...Q.."}, {"...Q..", "Q.....", "....Q.", ".Q....", ".....Q", "..Q..."}, {"....Q.", "..Q...", "Q.....", ".....Q", "...Q..", ".Q...."},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveNQueens(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}
