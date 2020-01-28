package p53_maximum_subarray

import (
	"github.com/TeslaCN/goleetcode/leetcode/p53_maximum-subarray/p53_1"
	"github.com/TeslaCN/goleetcode/leetcode/p53_maximum-subarray/p53_2"
	"testing"
)

type args struct {
	nums []int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"sample-0", args{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}}, 6},
	{"case-0", args{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 8}}, 9},
	{"case-1", args{[]int{-1, -2, -3, -1, -4, -1}}, -1},
	{"case-2", args{[]int{1, 2, 3, 4, 5, 6, 7}}, 28},
	{"case-3", args{[]int{-1}}, -1},
	{"case-4", args{[]int{-1, -1}}, -1},
}

func Test_1(t *testing.T) {
	testMaxSubArray(t, p53_1.MaxSubArray)
}
func Test_2(t *testing.T) {
	testMaxSubArray(t, p53_2.MaxSubArray)
}

func testMaxSubArray(t *testing.T, f func([]int) int) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.nums); got != tt.want {
				t.Errorf("MaxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
