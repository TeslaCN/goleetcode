package p617_2

import "github.com/TeslaCN/goleetcode/util"

type TreeNode = util.TreeNode

func MergeTrees(t1, t2 *TreeNode) *TreeNode {
	return mergeNode(t1, t2)
}

func mergeNode(n1, n2 *TreeNode) *TreeNode {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	return &TreeNode{
		Val:   n1.Val + n2.Val,
		Left:  mergeNode(n1.Left, n2.Left),
		Right: mergeNode(n1.Right, n2.Right),
	}
}
