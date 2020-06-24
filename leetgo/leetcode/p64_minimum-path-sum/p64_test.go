package p64_minimum_path_sum

import (
	"github.com/TeslaCN/goleetcode/leetcode/p64_minimum-path-sum/p64_1"
	"github.com/TeslaCN/goleetcode/leetcode/p64_minimum-path-sum/p64_2"
	"testing"
)

func Test_1(t *testing.T) {
	testMinPathSum(t, p64_1.MinPathSum)
}

func Test_2(t *testing.T) {
	testMinPathSum(t, p64_2.MinPathSum)
}

func testMinPathSum(t *testing.T, f func([][]int) int) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{[][]int{
			{1, 3, 1},
			{1, 5, 1},
			{4, 2, 1},
		}}, 7},
		{"sample-1", args{[][]int{
			{1, 2, 5},
			{3, 2, 1},
		}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.grid); got != tt.want {
				t.Errorf("minPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
