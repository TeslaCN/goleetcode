package p538_1

import "github.com/TeslaCN/goleetcode/util"

type TreeNode = util.TreeNode

func ConvertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	inorderTraversal(root, 0)
	return root
}

func inorderTraversal(node *TreeNode, add int) int {
	if node == nil {
		return 0
	}
	sum := inorderTraversal(node.Right, add)
	node.Val += sum + add
	return inorderTraversal(node.Left, node.Val) + node.Val - add
}
