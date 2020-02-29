package p215_kth_largest_element_in_an_array

import (
	"github.com/TeslaCN/goleetcode/leetcode/p215_kth-largest-element-in-an-array/p215_1"
	"github.com/TeslaCN/goleetcode/leetcode/p215_kth-largest-element-in-an-array/p215_2"
	"github.com/TeslaCN/goleetcode/leetcode/p215_kth-largest-element-in-an-array/p215_3"
	"testing"
)

func Test_1(t *testing.T) {
	testFindKthLargest(t, p215_1.FindKthLargest)
}

func Test_2(t *testing.T) {
	testFindKthLargest(t, p215_2.FindKthLargest)
}

func Test_3(t *testing.T) {
	testFindKthLargest(t, p215_3.FindKthLargest)
}

func testFindKthLargest(t *testing.T, f func([]int, int) int) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{[]int{3, 2, 1, 5, 6, 4}, 2}, 5},
		{"case-1", args{[]int{2, 1}, 2}, 1},
		{"case-8", args{[]int{3, 2, 1}, 2}, 2},
		{"case-9", args{[]int{4, 4, 3, 2, 1}, 3}, 3},
		{"case-3", args{[]int{2, 2, 1}, 2}, 2},
		{"sample-1", args{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4}, 4},
		{"case-0", args{[]int{2, 1, 3}, 2}, 2},
		{"case-6", args{[]int{2, 1, 3}, 3}, 1},
		{"case-7", args{[]int{2, 1, 3}, 1}, 3},
		{"case-2", args{[]int{1, 2}, 2}, 1},
		{"case-4", args{[]int{1, 1, 1, 1}, 2}, 1},
		{"case-5", args{[]int{5, 2, 3, 7, 6, 1, 4, 9, 8, 0}, 5}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
