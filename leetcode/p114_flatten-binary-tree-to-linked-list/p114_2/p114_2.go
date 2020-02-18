package p114_2

import "github.com/TeslaCN/goleetcode/util"

type TreeNode = util.TreeNode

var Flatten = flatten

// Space: O(1)
// Runtime: O(n)
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	right := root.Right
	root.Right = root.Left
	root.Left = nil
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
}
