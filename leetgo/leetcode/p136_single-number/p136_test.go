package p136_single_number

import (
	"github.com/TeslaCN/goleetcode/leetcode/p136_single-number/p136_1"
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
	{"sample-0", args{[]int{2, 2, 1}}, 1},
	{"sample-1", args{[]int{4, 1, 2, 1, 2}}, 4},
}

func TestSingleNumber(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p136_1.SingleNumber(tt.args.nums); got != tt.want {
				t.Errorf("SingleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
