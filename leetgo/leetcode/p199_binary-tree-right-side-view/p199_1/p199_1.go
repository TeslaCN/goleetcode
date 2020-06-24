package p199_1

import "github.com/TeslaCN/goleetcode/util"

type TreeNode = util.TreeNode

// Space: O(n)
// Runtime: O(n)
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	nodes, views := []*TreeNode{root}, make([]int, 0)
	cur, right := 0, 1
	for cur < len(nodes) {
		for ; cur < right; cur++ {
			if nodes[cur].Left != nil {
				nodes = append(nodes, nodes[cur].Left)
			}
			if nodes[cur].Right != nil {
				nodes = append(nodes, nodes[cur].Right)
			}
		}
		views = append(views, nodes[right-1].Val)
		right = len(nodes)
	}
	return views
}

/*
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 2.2 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
