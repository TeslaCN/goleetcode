package p94_1

import "github.com/TeslaCN/goleetcode/util"

type TreeNode = util.TreeNode

var InorderTraversal = inorderTraversal

// 递归
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	inorder(root, func(node *TreeNode) {
		result = append(result, node.Val)
	})
	return result
}

func inorder(node *TreeNode, f func(*TreeNode)) {
	if node == nil {
		return
	}
	inorder(node.Left, f)
	if f != nil {
		f(node)
	}
	inorder(node.Right, f)
}
