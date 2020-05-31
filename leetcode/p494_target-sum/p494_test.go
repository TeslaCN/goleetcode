package p494_target_sum

import (
	"github.com/TeslaCN/goleetcode/leetcode/p494_target-sum/p494_1"
	"github.com/TeslaCN/goleetcode/leetcode/p494_target-sum/p494_2"
	"github.com/TeslaCN/goleetcode/leetcode/p494_target-sum/p494_3"
	"testing"
)

func Test_1(t *testing.T) {
	testFindTargetSumWays(t, p494_1.FindTargetSumWays)
}

func Test_2(t *testing.T) {
	testFindTargetSumWays(t, p494_2.FindTargetSumWays)
}

func Test_3(t *testing.T) {
	testFindTargetSumWays(t, p494_3.FindTargetSumWays)
}

func testFindTargetSumWays(t *testing.T, f func([]int, int) int) {
	type args struct {
		nums []int
		S    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{[]int{1, 1, 1, 1, 1}, 3}, 5},
		{"sample-1", args{[]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1}, 256},
		{"case-0", args{[]int{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20}, 200}, 15504},
		//{"case-1", args{[]int{40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40}, 800}, 142506},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.nums, tt.args.S); got != tt.want {
				t.Errorf("findTargetSumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
