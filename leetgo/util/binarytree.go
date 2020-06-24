package util

import (
	"log"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateTreeFromSequenceInt(values []int) *TreeNode {
	return CreateTreeFromSequence(ConvertToIntPointer(values))
}

// Deprecated: Not suitable for leetcode
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

func PreorderTraversal(root *TreeNode, f func(*TreeNode)) {
	pre(root, f)
}

func pre(node *TreeNode, f func(*TreeNode)) {
	if node == nil {
		return
	}
	if f != nil {
		f(node)
	}
	pre(node.Left, f)
	pre(node.Right, f)
}

func InorderTraversal(root *TreeNode, f func(*TreeNode)) {
	inorder(root, f)
}
func inorder(node *TreeNode, f func(*TreeNode)) {
	if node == nil {
		return
	}
	inorder(node.Left, f)
	if f != nil {
		f(node)
	}
	inorder(node.Right, f)
}

func PostorderTraversal(root *TreeNode, f func(*TreeNode)) {
	postorder(root, f)
}
func postorder(node *TreeNode, f func(*TreeNode)) {
	if node == nil {
		return
	}
	postorder(node.Left, f)
	postorder(node.Right, f)
	if f != nil {
		f(node)
	}
}

// Leetcode string
// [4,-7,-3,null,null,-9,-3,9,-7,-4,null,6,null,-6,-6,null,null,0,6,5,null,9,null,null,-1,-4,null,null,null,-2]
func MakeTreeFromString(s string) *TreeNode {
	if s == "" {
		return nil
	}
	trim := strings.TrimSpace(strings.Trim(s, " []"))
	if trim == "" {
		return nil
	}
	split := strings.Split(trim, ",")
	values := make([]*int, len(split))
	for i, val := range split {
		val = strings.TrimSpace(val)
		if val == "null" {
			values[i] = nil
			continue
		}
		num, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalln(err)
		}
		values[i] = &num
	}
	return CreateTreeBySequenceTraversal(values)
}

func CreateTreeBySequenceTraversal(values []*int) *TreeNode {
	if len(values) < 1 || values[0] == nil {
		return nil
	}
	root := &TreeNode{Val: *values[0]}
	createNode([]*TreeNode{root}, values[1:])
	return root
}
func createNode(nodes []*TreeNode, values []*int) {
	cur := 0
	var nextLevel []*TreeNode
	for i := 0; i < len(nodes); i++ {
		if values[cur] != nil {
			nodes[i].Left = &TreeNode{Val: *values[cur]}
			nextLevel = append(nextLevel, nodes[i].Left)
		}
		cur++
		if cur == len(values) {
			break
		}
		if values[cur] != nil {
			nodes[i].Right = &TreeNode{Val: *values[cur]}
			nextLevel = append(nextLevel, nodes[i].Right)
		}
		cur++
		if cur == len(values) {
			break
		}
	}
	if cur < len(values) && len(nextLevel) > 0 {
		createNode(nextLevel, values[cur:])
	}
}
