package p236_3

import "github.com/TeslaCN/goleetcode/util"

type TreeNode = util.TreeNode

var LowestCommonAncestor = lowestCommonAncestor

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil || p.Val == root.Val || q.Val == root.Val {
		return root
	}
	left, right := lowestCommonAncestor(root.Left, p, q), lowestCommonAncestor(root.Right, p, q)
	switch {
	case left != nil && right != nil:
		return root
	case left != nil:
		return left
	case right != nil:
		return right
	}
	return nil
}

/*
执行用时 : 16 ms , 在所有 Go 提交中击败了 69.74% 的用户
内存消耗 : 7.1 MB , 在所有 Go 提交中击败了 50.00% 的用户
*/
