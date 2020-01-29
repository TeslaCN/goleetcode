package p101_1

import "github.com/TeslaCN/goleetcode/util"

type TreeNode = util.TreeNode

// 递归
func IsSymmetric(root *TreeNode) bool {
	return root == nil || isSubTreeSymmetric(root.Left, root.Right)
}

func isSubTreeSymmetric(left, right *TreeNode) bool {
	return left == nil && right == nil ||
		left != nil && right != nil && left.Val == right.Val &&
			isSubTreeSymmetric(left.Left, right.Right) &&
			isSubTreeSymmetric(left.Right, right.Left)
}
