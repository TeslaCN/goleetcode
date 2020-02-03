package p11_container_with_most_water

import (
	"github.com/TeslaCN/goleetcode/leetcode/p11_container-with-most-water/p11_1"
	"github.com/TeslaCN/goleetcode/leetcode/p11_container-with-most-water/p11_2"
	"testing"
)

type args struct {
	height []int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"sample-0", args{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, 49},
	{"sample-1", args{[]int{2, 3, 4, 5, 18, 17, 6}}, 17},
	{"case-0", args{[]int{1, 2, 3, 4, 5}}, 6},
	{"case-1", args{[]int{2, 3, 4, 5, 18, 100, 17, 6}}, 34},
	{"case-2", args{[]int{2, 3, 2, 1, 28, 27, 2, 3}}, 27},
}

func Test_1(t *testing.T) {
	testMaxArea(t, p11_1.MaxArea)
}

func Test_2(t *testing.T) {
	testMaxArea(t, p11_2.MaxArea)
}

func testMaxArea(t *testing.T, f func([]int) int) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
