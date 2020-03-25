package p892_surface_area_of_3d_shapes

import (
	"github.com/TeslaCN/goleetcode/leetcode/p892_surface-area-of-3d-shapes/p892_1"
	"github.com/TeslaCN/goleetcode/leetcode/p892_surface-area-of-3d-shapes/p892_2"
	"testing"
)

func Test_1(t *testing.T) {
	testSurfaceArea(t, p892_1.SurfaceArea)
}

func Test_2(t *testing.T) {
	testSurfaceArea(t, p892_2.SurfaceArea)
}

func testSurfaceArea(t *testing.T, f func([][]int) int) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{[][]int{{2}}}, 10},
		{"sample-1", args{[][]int{
			{1, 2},
			{3, 4},
		}}, 34},
		{"sample-2", args{[][]int{
			{1, 0},
			{0, 2},
		}}, 16},
		{"sample-3", args{[][]int{
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		}}, 32},
		{"sample-4", args{[][]int{
			{2, 2, 2},
			{2, 1, 2},
			{2, 2, 2},
		}}, 46},
		{"case-0", args{[][]int{
			{1, 0},
			{0, 1},
		}}, 12},
		{"case-1", args{[][]int{{1}}}, 6},
		{"case-2", args{[][]int{{3}}}, 14},
		{"case-3", args{[][]int{
			{1, 0, 1},
			{0, 0, 0},
			{1, 0, 1},
		}}, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.grid); got != tt.want {
				t.Errorf("surfaceArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
