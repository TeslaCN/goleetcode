package p455_assign_cookies

import (
	"github.com/TeslaCN/goleetcode/leetcode/p455_assign-cookies/p455_1"
	"github.com/TeslaCN/goleetcode/leetcode/p455_assign-cookies/p455_2"
	"testing"
)

func Test_1(t *testing.T) {
	testFindContentChildren(t, p455_1.FindContentChildren)
}

func Test_2(t *testing.T) {
	testFindContentChildren(t, p455_2.FindContentChildren)
}

func testFindContentChildren(t *testing.T, f func([]int, []int) int) {
	type args struct {
		g []int
		s []int
	}
	var tests = []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{[]int{1, 2, 3}, []int{1, 1}}, 1},
		{"sample-1", args{[]int{1, 2}, []int{1, 2, 3}}, 2},
		{"case-0", args{[]int{1, 2, 8, 12}, []int{1, 2, 3, 7, 11, 13}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.g, tt.args.s); got != tt.want {
				t.Errorf("findContentChildren() = %v, want %v", got, tt.want)
			}
		})
	}
}
