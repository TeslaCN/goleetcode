package p50_powx_n

import (
	"github.com/TeslaCN/goleetcode/leetcode/p50_powx-n/p50_1"
	"github.com/TeslaCN/goleetcode/leetcode/p50_powx-n/p50_2"
	"github.com/TeslaCN/goleetcode/leetcode/p50_powx-n/p50_3"
	"math"
	"testing"
)

func Test_1(t *testing.T) {
	testMyPow(t, p50_1.MyPow)
}

func Test_2(t *testing.T) {
	testMyPow(t, p50_2.MyPow)
}

func Test_3(t *testing.T) {
	testMyPow(t, p50_3.MyPow)
}

func testMyPow(t *testing.T, f func(float64, int) float64) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"sample-0", args{2.00000, 10}, 1024.00000},
		//{"sample-1", args{2.10000, 3}, 9.26100},
		{"sample-2", args{2.00000, -2}, 0.25000},
		{"case-0", args{2.00000, math.MinInt32}, 0},
		{"case-1", args{2.00000, math.MaxInt32}, math.Inf(1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
