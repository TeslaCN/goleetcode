package p167_two_sum_ii_input_array_is_sorted

import (
	"github.com/TeslaCN/goleetcode/leetcode/p167_two-sum-ii-input-array-is-sorted/p167_1"
	"reflect"
	"testing"
)

type args struct {
	numbers []int
	target  int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{"sample-0", args{numbers: []int{2, 7, 11, 15}, target: 9}, []int{1, 2}},
	{"case-0", args{numbers: []int{1, 2, 5, 7, 8, 11, 12, 17}, target: 16}, []int{3, 6}},
	{"case-1", args{numbers: []int{1, 2}, target: 3}, []int{1, 2}},
	{"case-2", args{numbers: []int{1, 2, 8, 9, 12}, target: 21}, []int{4, 5}},
}

func Test_twoSum(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p167_1.TwoSum(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
