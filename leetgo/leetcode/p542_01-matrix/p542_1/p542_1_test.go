package p542_1

import (
	"reflect"
	"testing"
)

func Test_updateMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"sample-0", args{[][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		}}, [][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		}},
		{"sample-1", args{[][]int{
			{0, 0, 0},
			{0, 1, 0},
			{1, 1, 1},
		}}, [][]int{
			{0, 0, 0},
			{0, 1, 0},
			{1, 2, 1},
		}},
		{"case-0", args{[][]int{
			{0, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
		}}, [][]int{
			{0, 1, 2},
			{1, 2, 3},
			{2, 3, 4},
		}},
		{"case-1", args{[][]int{
			{0, 1, 0},
			{1, 1, 1},
			{0, 1, 0},
		}}, [][]int{
			{0, 1, 0},
			{1, 2, 1},
			{0, 1, 0},
		}},
		{"case-2", args{[][]int{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		}}, [][]int{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateMatrix(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
