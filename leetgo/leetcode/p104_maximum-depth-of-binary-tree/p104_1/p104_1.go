package p104_1

import "github.com/TeslaCN/goleetcode/util"

type TreeNode = util.TreeNode

func MaxDepth(root *TreeNode) int {
	max, currentDepth := new(int), new(int)
	preorderTraversal(root, currentDepth, max)
	return *max
}

func preorderTraversal(node *TreeNode, currentDepth *int, max *int) {
	if node == nil {
		return
	}
	*currentDepth++
	if *currentDepth > *max {
		*max = *currentDepth
	}
	preorderTraversal(node.Left, currentDepth, max)
	preorderTraversal(node.Right, currentDepth, max)
	*currentDepth--
}
