package util

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateTreeFromSequence(values []*int) *TreeNode {
	if values == nil || len(values) < 1 {
		return nil
	}
	size := len(values)
	nodes := make([]*TreeNode, size)
	for i := 0; i < size; i++ {
		if values[i] != nil {
			nodes[i] = &TreeNode{Val: *values[i]}
		}
	}
	for i := 0; i < size/2; i++ {
		if nodes[i] != nil {
			var left, right int
			left = 2*i + 1
			right = 2*i + 2
			if left < size {
				nodes[i].Left = nodes[left]
			}
			if right < size {
				nodes[i].Right = nodes[right]
			}
		}
	}
	return nodes[0]
}

func SequenceTraversal(root *TreeNode) []*TreeNode {
	if root == nil {
		return make([]*TreeNode, 0)
	}
	var nodes []*TreeNode
	nodes = append(nodes, root)
	for i := 0; i < len(nodes); i++ {
		if nodes[i] != nil && (nodes[i].Left != nil || nodes[i].Right != nil) {
			nodes = append(nodes, nodes[i].Left)
			nodes = append(nodes, nodes[i].Right)
		}
	}
	return nodes
}

func SequenceTraversalValues(root *TreeNode) []*int {
	nodes := SequenceTraversal(root)
	values := make([]*int, len(nodes))
	for i := 0; i < len(nodes); i++ {
		if nodes[i] != nil {
			values[i] = &nodes[i].Val
		}
	}
	return values
}

func PreorderTraversal(root *TreeNode) {
	pre(root)
}

func pre(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Printf("%v ", node.Val)
	pre(node.Left)
	pre(node.Right)
}

func InorderTraversal(root *TreeNode) {
	inorder(root)
}
func inorder(node *TreeNode) {
	if node == nil {
		return
	}
	inorder(node.Left)
	fmt.Printf("%v ", node.Val)
	inorder(node.Right)
}

func PostorderTraversal(root *TreeNode) {
	postorder(root)
}
func postorder(node *TreeNode) {
	if node == nil {
		return
	}
	postorder(node.Left)
	postorder(node.Right)
	fmt.Printf("%v ", node.Val)
}
