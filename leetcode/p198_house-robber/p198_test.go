package p198_house_robber

import (
	"github.com/TeslaCN/goleetcode/leetcode/p198_house-robber/p198_1"
	"github.com/TeslaCN/goleetcode/leetcode/p198_house-robber/p198_2"
	"testing"
)

type args struct {
	nums []int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"sample-0", args{[]int{1, 2, 3, 1}}, 4},
	{"sample-1", args{[]int{2, 7, 9, 3, 1}}, 12},
	{"case-0", args{[]int{}}, 0},
	{"case-1", args{[]int{1}}, 1},
	{"case-2", args{[]int{1, 2, 3, 4}}, 6},
	{"case-3", args{[]int{4, 3, 2, 1}}, 6},
}

func Test_1(t *testing.T) {
	testRob(t, p198_1.Rob)
}

func Test_2(t *testing.T) {
	testRob(t, p198_2.Rob)
}

func testRob(t *testing.T, f func([]int) int) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.nums); got != tt.want {
				t.Errorf("Rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
