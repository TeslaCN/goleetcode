package p695_max_area_of_island

import (
	"github.com/TeslaCN/goleetcode/leetcode/p695_max-area-of-island/p695_1"
	"github.com/TeslaCN/goleetcode/leetcode/p695_max-area-of-island/p695_2"
	"github.com/TeslaCN/goleetcode/leetcode/p695_max-area-of-island/p695_3"
	"testing"
)

func Test_1(t *testing.T) {
	testMaxAreaOfIsland(t, p695_1.MaxAreaOfIsland)
}

func Test_2(t *testing.T) {
	testMaxAreaOfIsland(t, p695_2.MaxAreaOfIsland)
}

func Test_3(t *testing.T) {
	testMaxAreaOfIsland(t, p695_3.MaxAreaOfIsland)
}

func testMaxAreaOfIsland(t *testing.T, f func([][]int) int) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{[][]int{
			{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		}}, 6},
		{"sample-1", args{[][]int{{0, 0, 0, 0, 0, 0, 0, 0}}}, 0},
		{"case-0", args{[][]int{{}}}, 0},
		{"case-1", args{[][]int{{1}}}, 1},
		{"case-2", args{[][]int{{0}}}, 0},
		{"case-3", args{[][]int{{0, 0, 0, 1, 0, 1, 0}}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.grid); got != tt.want {
				t.Errorf("maxAreaOfIsland() = %v, want %v", got, tt.want)
			}
		})
	}
}
