package p96_unique_binary_search_trees

import (
	"github.com/TeslaCN/goleetcode/leetcode/p96_unique-binary-search-trees/p96_1"
	"github.com/TeslaCN/goleetcode/leetcode/p96_unique-binary-search-trees/p96_2"
	"testing"
)

// Deprecated
func Test_1(t *testing.T) {
	testNumTrees(t, p96_1.NumTrees)
}

func Test_2(t *testing.T) {
	testNumTrees(t, p96_2.NumTrees)
}

func testNumTrees(t *testing.T, f func(int) int) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{3}, 5},
		{"case-0", args{2}, 2},
		{"case-1", args{1}, 1},
		{"case-2", args{0}, 1},
		{"case-3", args{4}, 14},
		{"case-4", args{5}, 42},
		{"case-5", args{10}, 16796},
		{"case-6", args{15}, 9694845},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.n); got != tt.want {
				t.Errorf("numTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
