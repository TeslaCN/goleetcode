package p128_longest_consecutive_sequence

import (
	"github.com/TeslaCN/goleetcode/leetcode/p128_longest-consecutive-sequence/p128_1"
	"github.com/TeslaCN/goleetcode/leetcode/p128_longest-consecutive-sequence/p128_2"
	"testing"
)

func Test_1(t *testing.T) {
	testLongestConsecutive(t, p128_1.LongestConsecutive)
}

func Test_2(t *testing.T) {
	testLongestConsecutive(t, p128_2.LongestConsecutive)
}
func testLongestConsecutive(t *testing.T, f func([]int) int) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{[]int{100, 4, 200, 1, 3, 2}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.nums); got != tt.want {
				t.Errorf("longestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
