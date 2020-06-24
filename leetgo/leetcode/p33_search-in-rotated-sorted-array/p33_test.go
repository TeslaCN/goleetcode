package p33_search_in_rotated_sorted_array

import (
	"github.com/TeslaCN/goleetcode/leetcode/p33_search-in-rotated-sorted-array/p33_1"
	"github.com/TeslaCN/goleetcode/leetcode/p33_search-in-rotated-sorted-array/p33_2"
	"testing"
)

func Test_1(t *testing.T) {
	testSearch(t, p33_1.Search)
}

func Test_2(t *testing.T) {
	testSearch(t, p33_2.Search)
}

func testSearch(t *testing.T, f func([]int, int) int) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{[]int{4, 5, 6, 7, 0, 1, 2}, 0}, 4},
		{"sample-1", args{[]int{4, 5, 6, 7, 0, 1, 2}, 3}, -1},
		{"case-0", args{[]int{4, 5, 6, 0, 1, 2}, 3}, -1},
		{"case-1", args{[]int{4, 5, 6, 0, 1, 2}, 4}, 0},
		{"case-2", args{[]int{0, 1, 2, 3, 4, 5}, 4}, 4},
		{"case-3", args{[]int{0, 1, 2, 3, 4, 5, 6}, 4}, 4},
		{"case-4", args{[]int{1, 2, 3, 4, 5, 6, 0}, 4}, 3},
		{"case-5", args{[]int{1, 2, 3, 4, 6, 0}, 6}, 4},
		{"case-6", args{[]int{1, 2, 3, 4, 6, 0}, 0}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
