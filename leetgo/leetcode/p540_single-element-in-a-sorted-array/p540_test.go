package p540_single_element_in_a_sorted_array

import (
	"github.com/TeslaCN/goleetcode/leetcode/p540_single-element-in-a-sorted-array/p540_1"
	"github.com/TeslaCN/goleetcode/leetcode/p540_single-element-in-a-sorted-array/p540_2"
	"testing"
)

func Test_1(t *testing.T) {
	testSingleNonDuplicate(t, p540_1.SingleNonDuplicate)
}

func Test_2(t *testing.T) {
	testSingleNonDuplicate(t, p540_2.SingleNonDuplicate)
}

func testSingleNonDuplicate(t *testing.T, f func([]int) int) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{[]int{1, 1, 2, 3, 3, 4, 4, 8, 8}}, 2},
		{"case-7  ", args{[]int{1, 1, 2, 2, 3, 3, 4, 8, 8}}, 4},
		{"case-5  ", args{[]int{1, 1, 2, 3, 3, 4, 4}}, 2},
		{"case-6  ", args{[]int{1, 1, 2, 2, 3, 4, 4}}, 3},
		{"sample-1", args{[]int{3, 3, 7, 7, 10, 11, 11}}, 10},
		{"case-0  ", args{[]int{3, 7, 7, 10, 10, 11, 11}}, 3},
		{"case-1  ", args{[]int{3, 3, 7, 7, 10, 10, 11}}, 11},
		{"case-2  ", args{[]int{3, 3, 7, 7, 8, 10, 10, 11, 11}}, 8},
		{"case-3  ", args{[]int{1, 1, 2}}, 2},
		{"case-4  ", args{[]int{1, 2, 2}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.nums); got != tt.want {
				t.Errorf("singleNonDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p540_1.SingleNonDuplicate([]int{1, 1, 2, 2, 3, 3, 4, 8, 8})
	}
}
func Benchmark_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p540_2.SingleNonDuplicate([]int{1, 1, 2, 2, 3, 3, 4, 8, 8})
	}
}
