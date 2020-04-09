package p55_jump_game

import (
	"github.com/TeslaCN/goleetcode/leetcode/p55_jump-game/p55_1"
	"testing"
)

func Test_1(t *testing.T) {
	testCanJump(t, p55_1.CanJump)
}

func testCanJump(t *testing.T, f func([]int) bool) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"sample-0", args{[]int{2, 3, 1, 1, 4}}, true},
		{"sample-1", args{[]int{3, 2, 1, 0, 4}}, false},
		{"case-0", args{[]int{3, 2, 2, 0, 4}}, true},
		{"case-1", args{[]int{4, 0, 0, 0, 0}}, true},
		{"case-2", args{[]int{1, 1, 1, 1, 0}}, true},
		{"case-3", args{[]int{14, 0, 0, 0, 0}}, true},
		{"case-4", args{[]int{1, 1, 1, 9, 0}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
