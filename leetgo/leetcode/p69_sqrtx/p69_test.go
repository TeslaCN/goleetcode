package p69_sqrtx

import (
	"github.com/TeslaCN/goleetcode/leetcode/p69_sqrtx/p69_1"
	"github.com/TeslaCN/goleetcode/leetcode/p69_sqrtx/p69_2"
	"testing"
)

func Test_1(t *testing.T) {
	testMySqrt(t, p69_1.MySqrt)
}

func Test_2(t *testing.T) {
	testMySqrt(t, p69_2.MySqrt)
}

func testMySqrt(t *testing.T, f func(int) int) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{4}, 2},
		{"sample-1", args{8}, 2},
		{"case-0", args{360}, 18},
		{"case-1", args{81}, 9},
		{"case-2", args{80}, 8},
		{"case-3", args{0}, 0},
		{"case-4", args{1}, 1},
		{"case-5", args{2}, 1},
		{"case-6", args{3}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
