package p236_2

import "github.com/TeslaCN/goleetcode/util"

type TreeNode = util.TreeNode

var LowestCommonAncestor = lowestCommonAncestor

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil || root.Val == q.Val || root.Val == p.Val {
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
	for pIndex < rootIndex && qIndex < rootIndex || pIndex > rootIndex && qIndex > rootIndex {
		if pIndex < rootIndex && qIndex < rootIndex {
			root = root.Left
		} else if pIndex > rootIndex && qIndex > rootIndex {
			root = root.Right
		}
		for i, val := range inorder {
			if root.Val == val {
				rootIndex = i
				break
			}
		}
	}
	switch root.Val {
	case p.Val:
		return p
	case q.Val:
		return q
	default:
		return root
	}
}

func inorderTraversal(node *TreeNode, f func(*TreeNode)) {
	if node != nil {
		inorderTraversal(node.Left, f)
		f(node)
		inorderTraversal(node.Right, f)
	}
}

/*
执行用时 : 64 ms , 在所有 Go 提交中击败了 9.78% 的用户
内存消耗 : 6.6 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
