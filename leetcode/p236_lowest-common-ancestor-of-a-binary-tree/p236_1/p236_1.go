package p236_1

import "github.com/TeslaCN/goleetcode/util"

type TreeNode = util.TreeNode

// Deprecated: TLE
var LowestCommonAncestor = lowestCommonAncestor

// Deprecated: TLE
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}
	inorder := make([]int, 0)
	rootIndex, pIndex, qIndex := -1, -1, -1
	inorderTraversal(root, func(node *TreeNode) {
		switch node.Val {
		case root.Val:
			rootIndex = len(inorder)
		case p.Val:
			pIndex = len(inorder)
		case q.Val:
			qIndex = len(inorder)
		}
		inorder = append(inorder, node.Val)
	})
	if pIndex < rootIndex && qIndex > rootIndex || pIndex > rootIndex && qIndex < rootIndex {
		return root
	} else if pIndex < rootIndex && qIndex < rootIndex {
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		return lowestCommonAncestor(root.Right, p, q)
	}
}

func inorderTraversal(node *TreeNode, f func(*TreeNode)) {
	if node != nil {
		inorderTraversal(node.Left, f)
		f(node)
		inorderTraversal(node.Right, f)
	}
}
