package util

import (
	"reflect"
	"testing"
)

/*
	     3
	    / \
	   4   5
	  / \   \
	 5   4   7

	0  1  2  3  4   5   6
	3, 4, 5, 5, 4, nil, 7
	0  1  1  2  2   2   2
	2n+1
*/
var root = &TreeNode{
	Val: 3,
	Left: &TreeNode{
		Val:   4,
		Left:  &TreeNode{Val: 5},
		Right: &TreeNode{Val: 4},
	},
	Right: &TreeNode{
		Val:   5,
		Right: &TreeNode{Val: 7},
	},
}

func TestSequenceTraversal(t *testing.T) {
	nodes := SequenceTraversal(root)
	t.Logf("%v", nodes)
	values := SequenceTraversalValues(root)
	t.Log(FormatIntPointerSlice(values))
}

func TestCreateTreeFromSequence(t *testing.T) {
	//nums := []int{3, 4, 5, 5, 4, -1, 7}
	nums := []int{1, -1, 2, -1, 3}
	size := len(nums)
	inputs := make([]*int, size)
	for i := 0; i < size; i++ {
		if nums[i] != -1 {
			inputs[i] = &nums[i]
		}
	}
	tree := CreateTreeFromSequence(inputs)
	t.Logf("DeepEqual: %v", reflect.DeepEqual(root, tree))
}

func TestPreorderTraversal(t *testing.T) {
	PreorderTraversal(root)
}

func TestInorderTraversal(t *testing.T) {
	InorderTraversal(root)
}

func TestPostorderTraversal(t *testing.T) {
	PostorderTraversal(root)
}
