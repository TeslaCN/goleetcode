package p283_move_zeroes

import (
	"github.com/TeslaCN/goleetcode/leetcode/p283_move-zeroes/p283_1"
	"reflect"
	"testing"
)

type args struct {
	nums []int
}

var tests = []struct {
	name   string
	args   args
	expect []int
}{
	{"sample-0", args{[]int{0, 1, 0, 3, 12}}, []int{1, 3, 12, 0, 0}},
	{"case-0", args{[]int{}}, []int{}},
	{"case-1", args{[]int{0}}, []int{0}},
	{"case-2", args{[]int{0, 0}}, []int{0, 0}},
	{"case-3", args{[]int{0, 1}}, []int{1, 0}},
}

func TestMoveZeroes(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p283_1.MoveZeroes(tt.args.nums)
			if !reflect.DeepEqual(tt.expect, tt.args.nums) {
				t.Errorf("Expect: %v, Got: %v", tt.expect, tt.args.nums)
			}
		})
	}
}
