package p581_shortest_unsorted_continuous_subarray

import (
	"github.com/TeslaCN/goleetcode/leetcode/p581_shortest-unsorted-continuous-subarray/p581_1"
	"github.com/TeslaCN/goleetcode/leetcode/p581_shortest-unsorted-continuous-subarray/p581_2"
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
	{"sample-0", args{[]int{2, 6, 4, 8, 10, 9, 15}}, 5},
	{"sample-1", args{[]int{2, 3, 3, 2, 4}}, 3},
	{"sample-2", args{[]int{1, 2, 4, 5, 3}}, 3},
	{"sample-3", args{[]int{1, 3, 5, 2, 4}}, 4},
	{"sample-4", args{[]int{2, 1, 5, 3, 4}}, 5},
	{"case-0", args{[]int{1, 2, 3, 4, 5, 4}}, 2},
	{"case-1", args{[]int{1, 2, 3, 6, 5, 4}}, 3},
	{"case-2", args{[]int{1, 2, 3, 4, 5, 6}}, 0},
	{"case-3", args{[]int{6, 2, 3, 4, 5, 6}}, 5},
	{"case-4", args{[]int{1, 1, 1, 1, 1, 1}}, 0},
	{"case-5", args{[]int{1, 1, 2, 1, 1, 1}}, 4},
	{"case-6", args{[]int{1, 2, 4, 5, 3, 9}}, 3},
	{"case-7", args{[]int{1, 2, 4, 3, 7}}, 2},
	{"case-8", args{[]int{2, 3, 2, 3, 4}}, 2},
	{"case-9", args{[]int{2, 3, 3, 2, 3}}, 3},
	{"case-10", args{[]int{2, 3, 3, 2, 2, 3, 3}}, 4},
	{"case-11", args{[]int{4, 3, 2, 1}}, 4},
	{"case-12", args{[]int{4, 3, 5, 2, 1}}, 5},
}

func Test_1(t *testing.T) {
	testFindUnsortedSubarray(t, p581_1.FindUnsortedSubarray)
}

func Test_2(t *testing.T) {
	testFindUnsortedSubarray(t, p581_2.FindUnsortedSubarray)
}

func testFindUnsortedSubarray(t *testing.T, f func([]int) int) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.nums); got != tt.want {
				t.Errorf("findUnsortedSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
