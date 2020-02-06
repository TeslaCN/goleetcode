package p94_2

import (
	"container/list"
	"github.com/TeslaCN/goleetcode/util"
)

type TreeNode = util.TreeNode

var InorderTraversal = inorderTraversal

// 迭代
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)
	stack := list.New()
	node := root
	for node != nil || stack.Len() > 0 {
		for node != nil {
			stack.PushBack(node)
			node = node.Left
		}
		node = stack.Remove(stack.Back()).(*TreeNode)
		result = append(result, node.Val)
		node = node.Right
	}
	return result
}
