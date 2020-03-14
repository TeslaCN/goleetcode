package p300_longest_increasing_subsequence

import (
	"github.com/TeslaCN/goleetcode/leetcode/p300_longest-increasing-subsequence/p300_1"
	"github.com/TeslaCN/goleetcode/leetcode/p300_longest-increasing-subsequence/p300_2"
	"testing"
)

func Test_1(t *testing.T) {
	testLengthOfLIS(t, p300_1.LengthOfLIS)
}

func Test_2(t *testing.T) {
	testLengthOfLIS(t, p300_2.LengthOfLIS)
}

func testLengthOfLIS(t *testing.T, f func([]int) int) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{[]int{10, 9, 2, 5, 3, 7, 101, 18}}, 4},
		{"case-0", args{[]int{10, 9, 2, 5, 3, 4, 101, 18}}, 4},
		{"case-1", args{[]int{8, 10, 9, 2, 5, 3, 4, 101, 18, 19}}, 5},
		{"case-2", args{[]int{8, 9, 10, 2, 5, 3, 4, 101, 18, 19}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
