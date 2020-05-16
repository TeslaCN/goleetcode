package p102_2

import "github.com/TeslaCN/goleetcode/util"

type TreeNode = util.TreeNode

var LevelOrder = levelOrder

// Runtime: O(n)
// Space: O(n)
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	levels, nodes := make([][]int, 0, 1), []*TreeNode{root}
	to := len(nodes)
	for from := 0; from < len(nodes); {
		level := make([]int, 0)
		for i := from; i < to; i++ {
			level = append(level, nodes[i].Val)
			if nodes[i].Left != nil {
				nodes = append(nodes, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				nodes = append(nodes, nodes[i].Right)
			}
		}
		levels = append(levels, level)
		from, to = to, len(nodes)
	}
	return levels
}

/*
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 2.8 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
