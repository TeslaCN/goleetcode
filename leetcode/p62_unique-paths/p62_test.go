package p62_unique_paths

import (
	"github.com/TeslaCN/goleetcode/leetcode/p62_unique-paths/p62_1"
	"github.com/TeslaCN/goleetcode/leetcode/p62_unique-paths/p62_2"
	"testing"
)

func Test_1(t *testing.T) {
	testUniquePaths(t, p62_1.UniquePaths)
}

func Test_2(t *testing.T) {
	testUniquePaths(t, p62_2.UniquePaths)
}

func testUniquePaths(t *testing.T, f func(int, int) int) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{3, 2}, 3},
		{"sample-1", args{7, 3}, 28},
		{"case-0", args{1, 1}, 1},
		//{"case-1", args{0, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
