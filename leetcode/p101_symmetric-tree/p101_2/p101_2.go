package p101_2

import (
	"container/list"
	"github.com/TeslaCN/goleetcode/util"
)

type TreeNode = util.TreeNode

// 迭代
func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var left, right *TreeNode
	stack := list.New()
	stack.PushBack(root.Left)
	stack.PushBack(root.Right)
	for stack.Len() > 0 {
		right = stack.Remove(stack.Back()).(*TreeNode)
		left = stack.Remove(stack.Back()).(*TreeNode)
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		stack.PushBack(left.Left)
		stack.PushBack(right.Right)
		stack.PushBack(left.Right)
		stack.PushBack(right.Left)
	}
	return true
}
