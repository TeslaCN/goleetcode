package lcof_64_qiu_12n_lcof

import (
	"github.com/TeslaCN/goleetcode/leetcode/lcof_64_qiu-12n-lcof/lcof_64_1"
	"github.com/TeslaCN/goleetcode/leetcode/lcof_64_qiu-12n-lcof/lcof_64_2"
	"testing"
)

func Test_1(t *testing.T) {
	testSumNums(t, lcof_64_1.SumNums)
}

func Test_2(t *testing.T) {
	testSumNums(t, lcof_64_2.SumNums)
}

func testSumNums(t *testing.T, f func(int) int) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{3}, 6},
		{"sample-1", args{9}, 45},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.n); got != tt.want {
				t.Errorf("sumNums() = %v, want %v", got, tt.want)
			}
		})
	}
}
