package p572_1

import "github.com/TeslaCN/goleetcode/util"

type TreeNode = util.TreeNode

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil || t == nil {
		return false
	}
	return equals(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func equals(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	return s.Val == t.Val && equals(s.Left, t.Left) && equals(s.Right, t.Right)
}

/*
执行用时 : 20 ms , 在所有 Go 提交中击败了 93.51% 的用户
内存消耗 : 6.2 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
