package p105_1

import "github.com/TeslaCN/goleetcode/util"

type TreeNode = util.TreeNode

func buildTree(preorder []int, inorder []int) *TreeNode {
	switch len(preorder) {
	case 0:
		return nil
	case 1:
		return &TreeNode{Val: preorder[0]}
	}
	root := &TreeNode{Val: preorder[0]}
	mid := 0
	for ; inorder[mid] != preorder[0]; mid++ {
	}
	root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])
	return root
}
