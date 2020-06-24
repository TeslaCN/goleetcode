package p102_1

import "github.com/TeslaCN/goleetcode/util"

type TreeNode = util.TreeNode

var LevelOrder = levelOrder

// Runtime: O(n)
// Space: O(n)
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	return traverseLevel([]*TreeNode{root}, make([][]int, 0))
}

func traverseLevel(nodes []*TreeNode, result [][]int) [][]int {
	if len(nodes) < 1 {
		return result
	}
	var nextLevel []*TreeNode
	var level []int
	for _, node := range nodes {
		level = append(level, node.Val)
		if node.Left != nil {
			nextLevel = append(nextLevel, node.Left)
		}
		if node.Right != nil {
			nextLevel = append(nextLevel, node.Right)
		}
	}
	return traverseLevel(nextLevel, append(result, level))
}
