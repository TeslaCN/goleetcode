package p226_1

import "github.com/TeslaCN/goleetcode/util"

type TreeNode = util.TreeNode

func InvertTree(root *TreeNode) *TreeNode {
	return invertNode(root)
}

func invertNode(node *TreeNode) *TreeNode {
	if node == nil {
		return node
	}
	node.Left, node.Right = invertNode(node.Right), invertNode(node.Left)
	return node
}
