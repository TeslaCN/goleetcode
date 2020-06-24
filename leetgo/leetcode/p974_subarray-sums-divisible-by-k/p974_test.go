package p974_subarray_sums_divisible_by_k

import (
	"github.com/TeslaCN/goleetcode/leetcode/p974_subarray-sums-divisible-by-k/p974_2"
	"github.com/TeslaCN/goleetcode/leetcode/p974_subarray-sums-divisible-by-k/p974_3"
	"testing"
)

//func Test_1(t *testing.T) {
//	testSubarraysDivByK(t, p974_1.SubarraysDivByK)
//}

func Test_2(t *testing.T) {
	testSubarraysDivByK(t, p974_2.SubarraysDivByK)
}

func Test_3(t *testing.T) {
	testSubarraysDivByK(t, p974_3.SubarraysDivByK)
}

func testSubarraysDivByK(t *testing.T, f func([]int, int) int) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{[]int{4, 5, 0, -2, -3, 1}, 5}, 7},
		{"sample-1", args{[]int{5}, 9}, 0},
		{"sample-2", args{[]int{-1, 2, 9}, 2}, 2},
		{"case-0", args{[]int{1, 2, 9}, 2}, 2},
		{"case-30000", args{make([]int, 30000), 10000}, 450015000},
		{"case-100000", args{make([]int, 100000), 10000}, 5000050000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.A, tt.args.K); got != tt.want {
				t.Errorf("subarraysDivByK() = %v, want %v", got, tt.want)
			}
		})
	}
}
