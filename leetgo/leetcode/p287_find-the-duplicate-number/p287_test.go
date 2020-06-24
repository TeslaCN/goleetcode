package p287_find_the_duplicate_number

import (
	"github.com/TeslaCN/goleetcode/leetcode/p287_find-the-duplicate-number/p287_1"
	"github.com/TeslaCN/goleetcode/leetcode/p287_find-the-duplicate-number/p287_2"
	"testing"
)

func Test_1(t *testing.T) {
	testFindDuplicate(t, p287_1.FindDuplicate)
}

func Test_2(t *testing.T) {
	testFindDuplicate(t, p287_2.FindDuplicate)
}

func testFindDuplicate(t *testing.T, f func([]int) int) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{[]int{1, 3, 4, 2, 2}}, 2},
		{"sample-1", args{[]int{3, 1, 3, 4, 2}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.nums); got != tt.want {
				t.Errorf("findDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
