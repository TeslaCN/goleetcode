package p98_1

import "github.com/TeslaCN/goleetcode/util"

type TreeNode = util.TreeNode

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	values := make([]int, 0)
	inOrderTraversal(root, func(node *TreeNode) {
		values = append(values, node.Val)
	})
	for i := 1; i < len(values); i++ {
		if values[i-1] >= values[i] {
			return false
		}
	}
	return true
}

func inOrderTraversal(node *TreeNode, f func(*TreeNode)) {
	if node == nil {
		return
	}
	inOrderTraversal(node.Left, f)
	f(node)
	inOrderTraversal(node.Right, f)
}

/*
执行用时 : 4 ms , 在所有 Go 提交中击败了 98.63% 的用户
内存消耗 : 5.8 MB , 在所有 Go 提交中击败了 40.00% 的用户
*/
