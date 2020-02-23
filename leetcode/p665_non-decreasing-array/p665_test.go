package p665_non_decreasing_array

import (
	"github.com/TeslaCN/goleetcode/leetcode/p665_non-decreasing-array/p665_1"
	"github.com/TeslaCN/goleetcode/leetcode/p665_non-decreasing-array/p665_2"
	"testing"
)

func Test_1(t *testing.T) {
	testCheckPossibility(t, p665_1.CheckPossibility)
}

func Test_2(t *testing.T) {
	testCheckPossibility(t, p665_2.CheckPossibility)
}

func testCheckPossibility(t *testing.T, f func([]int) bool) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"sample-0", args{[]int{4, 2, 3}}, true},
		{"sample-1", args{[]int{4, 2, 1}}, false},
		{"sample-2", args{[]int{3, 4, 2, 3}}, false},
		{"sample-3", args{[]int{2, 3, 3, 2, 4}}, true},
		{"case-0", args{[]int{2, 2, 3, 2, 2, 4}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.nums); got != tt.want {
				t.Errorf("checkPossibility() = %v, want %v", got, tt.want)
			}
		})
	}
}
