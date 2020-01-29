package p101_symmetric_tree

import (
	"github.com/TeslaCN/goleetcode/leetcode/p101_symmetric-tree/p101_1"
	"github.com/TeslaCN/goleetcode/leetcode/p101_symmetric-tree/p101_2"
	"github.com/TeslaCN/goleetcode/util"
	"testing"
)

type TreeNode = util.TreeNode

type args struct {
	root *TreeNode
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{"sample-0", args{util.CreateTreeFromSequenceInt([]int{1, 2, 2, 3, 4, 4, 3})}, true},
	{"sample-1", args{util.CreateTreeFromSequenceInt([]int{1, 2, 2, 0, 3, 0, 3})}, false},
	{"case-0", args{util.CreateTreeFromSequenceInt([]int{1})}, true},
	{"case-1", args{util.CreateTreeFromSequenceInt([]int{1, 2, 0})}, false},
	{"case-2", args{}, true},
}

func Test_1(t *testing.T) {
	testIsSymmetric(t, p101_1.IsSymmetric)
}

func Test_2(t *testing.T) {
	testIsSymmetric(t, p101_2.IsSymmetric)
}

func testIsSymmetric(t *testing.T, f func(*TreeNode) bool) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.root); got != tt.want {
				t.Errorf("IsSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
