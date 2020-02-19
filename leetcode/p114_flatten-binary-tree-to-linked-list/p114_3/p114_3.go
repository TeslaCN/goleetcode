package p114_3

import "github.com/TeslaCN/goleetcode/util"

type TreeNode = util.TreeNode

var Flatten = flatten

// Space: O(1)
// Runtime: O(n)
// iteration
func flatten(root *TreeNode) {
	p := root
	for p != nil {
		if p.Left == nil {
			p = p.Right
			continue
		}
		right := p.Right
		p.Right = p.Left
		p.Left = nil
		t := p
		for t.Right != nil {
			t = t.Right
		}
		t.Right = right
		p = p.Right
	}
}
