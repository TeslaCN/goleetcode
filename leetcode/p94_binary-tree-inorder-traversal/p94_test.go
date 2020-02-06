package p94_binary_tree_inorder_traversal

import (
	"github.com/TeslaCN/goleetcode/leetcode/p94_binary-tree-inorder-traversal/p94_1"
	"github.com/TeslaCN/goleetcode/leetcode/p94_binary-tree-inorder-traversal/p94_2"
	"github.com/TeslaCN/goleetcode/util"
	"reflect"
	"testing"
)

type TreeNode = util.TreeNode

type args struct {
	root *TreeNode
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{"sample-0", args{
		root: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 3},
			},
		},
	}, []int{1, 3, 2}},
	{
		/*
		     3
		    / \
		   1   4
		    \   \
		     2   5
		*/
		"case-0", args{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: &TreeNode{Val: 2},
				},
				Right: &TreeNode{
					Val:   4,
					Right: &TreeNode{Val: 5},
				},
			},
		},
		[]int{1, 2, 3, 4, 5},
	},
}

func Test_1(t *testing.T) {
	testInorderTraversal(t, p94_1.InorderTraversal)
}

func Test_2(t *testing.T) {
	testInorderTraversal(t, p94_2.InorderTraversal)
}

func testInorderTraversal(t *testing.T, f func(*TreeNode) []int) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
