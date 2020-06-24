package p977_squares_of_a_sorted_array

import (
	"github.com/TeslaCN/goleetcode/leetcode/p977_squares-of-a-sorted-array/p977_1"
	"github.com/TeslaCN/goleetcode/leetcode/p977_squares-of-a-sorted-array/p977_2"
	"github.com/TeslaCN/goleetcode/leetcode/p977_squares-of-a-sorted-array/p977_3"
	"reflect"
	"testing"
)

func Test_1(t *testing.T) {
	testSortedSquares(t, p977_1.SortedSquares)
}

func Test_2(t *testing.T) {
	testSortedSquares(t, p977_2.SortedSquares)
}

func Test_3(t *testing.T) {
	testSortedSquares(t, p977_3.SortedSquares)
}

func testSortedSquares(t *testing.T, f func([]int) []int) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"sample-0", args{[]int{-4, -1, 0, 3, 10}}, []int{0, 1, 9, 16, 100}},
		{"sample-1", args{[]int{-7, -3, 2, 3, 11}}, []int{4, 9, 9, 49, 121}},
		{"sample-2", args{[]int{-2, -1, 3}}, []int{1, 4, 9}},
		{"case-0", args{[]int{1, 2, 3, 4, 5}}, []int{1, 4, 9, 16, 25}},
		{"case-1", args{[]int{-5, -4, -3, -2, -1}}, []int{1, 4, 9, 16, 25}},
		{"case-2", args{[]int{-2, -1, 0, 1, 2}}, []int{0, 1, 1, 4, 4}},
		{"case-3", args{[]int{-5, -4, -3, -2, -1, 0, 1, 2}}, []int{0, 1, 1, 4, 4, 9, 16, 25}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
