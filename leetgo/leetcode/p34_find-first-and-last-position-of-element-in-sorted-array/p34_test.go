package p34_find_first_and_last_position_of_element_in_sorted_array

import (
	"github.com/TeslaCN/goleetcode/leetcode/p34_find-first-and-last-position-of-element-in-sorted-array/p34_1"
	"reflect"
	"testing"
)

type args struct {
	nums   []int
	target int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{"sample-0", args{[]int{5, 7, 7, 8, 8, 10}, 8}, []int{3, 4}},
	{"sample-1", args{[]int{5, 7, 7, 8, 8, 10}, 6}, []int{-1, -1}},
	{"case-0", args{[]int{5, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 8, 8, 10}, 7}, []int{1, 11}},
	{"case-1", args{[]int{5, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 8, 8, 10}, 8}, []int{12, 13}},
	{"case-2", args{[]int{1, 2, 3}, 2}, []int{1, 1}},
	{"case-3", args{[]int{1, 1, 2, 3}, 1}, []int{0, 1}},
	{"case-4", args{[]int{1}, 1}, []int{0, 0}},
	{"case-5", args{[]int{3, 3, 3}, 3}, []int{0, 2}},
}

func Test_searchRange(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p34_1.SearchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
