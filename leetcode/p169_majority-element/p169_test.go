package p169_majority_element

import (
	"github.com/TeslaCN/goleetcode/leetcode/p169_majority-element/p169_1"
	"github.com/TeslaCN/goleetcode/leetcode/p169_majority-element/p169_2"
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
	{"sample-0", args{[]int{3, 2, 3}}, 3},
	{"sample-1", args{[]int{2, 2, 1, 1, 1, 2, 2}}, 2},
	{"case-0", args{[]int{2, 1, 2, 1, 2, 2, 2, 3, 3}}, 2},
}

func Test_1(t *testing.T) {
	testMajorityElement(t, p169_1.MajorityElement)
}

func Test_2(t *testing.T) {
	testMajorityElement(t, p169_2.MajorityElement)
}

func testMajorityElement(t *testing.T, f func([]int) int) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.nums); got != tt.want {
				t.Errorf("MajorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
