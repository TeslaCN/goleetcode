package p114_1

import "github.com/TeslaCN/goleetcode/util"

type TreeNode = util.TreeNode

var Flatten = flatten

func flatten(root *TreeNode) {
	var nodes []*TreeNode
	preorder(root, func(node *TreeNode) {
		nodes = append(nodes, node)
	})
	for i := 0; i < len(nodes); i++ {
		nodes[i].Left = nil
		if i == len(nodes)-1 {
			nodes[i].Right = nil
		} else {
			nodes[i].Right = nodes[i+1]
		}
	}
}

func preorder(root *TreeNode, f func(node *TreeNode)) {
	if root == nil {
		return
	}
	f(root)
	preorder(root.Left, f)
	preorder(root.Right, f)
}
