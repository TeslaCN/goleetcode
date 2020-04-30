package p560_subarray_sum_equals_k

import (
	"github.com/TeslaCN/goleetcode/leetcode/p560_subarray-sum-equals-k/p560_1"
	"github.com/TeslaCN/goleetcode/leetcode/p560_subarray-sum-equals-k/p560_2"
	"testing"
)

func Test_1(t *testing.T) {
	testSubarraySum(t, p560_1.SubarraySum)
}

func Test_2(t *testing.T) {
	testSubarraySum(t, p560_2.SubarraySum)
}

func testSubarraySum(t *testing.T, f func([]int, int) int) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{[]int{1, 1, 1}, 2}, 2},
		{"case-0", args{[]int{-1, -1, -1, -3, 2, 1, -2, -1}, -3}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("subarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
