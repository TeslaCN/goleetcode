package p70_climbing_stairs

import (
	"github.com/TeslaCN/goleetcode/leetcode/p70_climbing-stairs/p70_1"
	"github.com/TeslaCN/goleetcode/leetcode/p70_climbing-stairs/p70_2"
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
	{"sample-0", args{2}, 2},
	{"sample-1", args{3}, 3},
	{"custom-0", args{4}, 5},
	{"custom-1", args{10}, 89},
	{"custom-2", args{44}, 1134903170},
}

func Test_1(t *testing.T) {
	testClimbStairs(t, p70_1.ClimbStairs)
}

func Test_2(t *testing.T) {
	testClimbStairs(t, p70_2.ClimbStairs)
}

func testClimbStairs(t *testing.T, f func(int) int) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
