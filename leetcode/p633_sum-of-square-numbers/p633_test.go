package p633_sum_of_square_numbers

import (
	"github.com/TeslaCN/goleetcode/leetcode/p633_sum-of-square-numbers/p633_1"
	"testing"
)

type args struct {
	c int
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{"sample-0", args{5}, true},
	{"sample-1", args{3}, false},
	{"case-0", args{25}, true},
	{"case-1", args{0}, true},
}

func Test_judgeSquareSum(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p633_1.JudgeSquareSum(tt.args.c); got != tt.want {
				t.Errorf("judgeSquareSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
