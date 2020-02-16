package p104_maximum_depth_of_binary_tree

import (
	"github.com/TeslaCN/goleetcode/leetcode/p104_maximum-depth-of-binary-tree/p104_1"
	"github.com/TeslaCN/goleetcode/util"
	"math"
	"testing"
)

type TreeNode = util.TreeNode

type args struct {
	root *TreeNode
}

var tests []struct {
	name string
	args args
	want int
}

func init() {
	tests = append(tests, []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{util.CreateTreeFromSequence(util.ConvertToIntPointer([]int{3, 9, 20, math.MinInt32, math.MinInt32, 15, 7}))}, 3},
		{"sample-1", args{util.CreateTreeFromSequence(util.ConvertToIntPointer([]int{}))}, 0},
	}...)
}

func TestMaxDepth(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p104_1.MaxDepth(tt.args.root); got != tt.want {
				t.Errorf("MaxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
