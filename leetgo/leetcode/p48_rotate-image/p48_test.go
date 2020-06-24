package p48_rotate_image

import (
	"github.com/TeslaCN/goleetcode/leetcode/p48_rotate-image/p48_1"
	"github.com/TeslaCN/goleetcode/leetcode/p48_rotate-image/p48_2"
	"reflect"
	"testing"
)

func Test_1(t *testing.T) {
	testRotate(t, p48_1.Rotate)
}
func Test_2(t *testing.T) {
	testRotate(t, p48_2.Rotate)
}
func testRotate(t *testing.T, f func([][]int)) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"case-0", args{[][]int{
			{1, 2},
			{3, 4},
		}}, [][]int{
			{3, 1},
			{4, 2},
		}},
		{"sample-0", args{[][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}}, [][]int{
			{7, 4, 1},
			{8, 5, 2},
			{9, 6, 3},
		}},
		{"sample-1", args{[][]int{
			{5, 1, 9, 11},
			{2, 4, 8, 10},
			{13, 3, 6, 7},
			{15, 14, 12, 16},
		}}, [][]int{
			{15, 13, 2, 5},
			{14, 3, 4, 1},
			{12, 6, 8, 9},
			{16, 7, 10, 11},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if f(tt.args.matrix); !reflect.DeepEqual(tt.args.matrix, tt.want) {
				t.Errorf("rotate() = %v, want %v", tt.args.matrix, tt.want)
			}
		})
	}
}
