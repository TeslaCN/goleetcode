package p42_trapping_rain_water

import (
	"github.com/TeslaCN/goleetcode/leetcode/p42_trapping-rain-water/p42_1"
	"github.com/TeslaCN/goleetcode/leetcode/p42_trapping-rain-water/p42_2"
	"github.com/TeslaCN/goleetcode/leetcode/p42_trapping-rain-water/p42_3"
	"testing"
)

func Test_1(t *testing.T) {
	testTrap(t, p42_1.Trap)
}

func Test_2(t *testing.T) {
	testTrap(t, p42_2.Trap)
}

func Test_3(t *testing.T) {
	testTrap(t, p42_3.Trap)
}

func testTrap(t *testing.T, f func([]int) int) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}}, 6},
		{"sample-1", args{[]int{5, 4, 1, 2}}, 1},
		{"sample-2", args{[]int{}}, 0},
		{"case-0", args{[]int{0, 1, 0, 1, 0, 1, 0}}, 2},
		{"case-1", args{[]int{0, 1, 1, 1, 1, 1, 0}}, 0},
		{"case-2", args{[]int{0, 3, 1, 1, 2, 1, 0}}, 2},
		{"case-3", args{[]int{0, 3, 3, 5, 5, 6, 5}}, 0},
		{"case-4", args{[]int{0, 0, 0, 0, 0}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.height); got != tt.want {
				t.Errorf("trap() = %v, want %v", got, tt.want)
			}
		})
	}
}
