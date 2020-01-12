package p70

import (
	use "goleetcode/p70/p70_2"
	"testing"
)

type args struct {
	n int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"sample", args{2}, 2},
	{"sample", args{3}, 3},
	{"custom", args{4}, 5},
	{"custom", args{10}, 89},
	{"custom", args{44}, 1134903170},
}

func Test_climbStairs1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := use.ClimbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
