package p543_1

import "github.com/TeslaCN/goleetcode/util"

type TreeNode = util.TreeNode

var DiameterOfBinaryTree = diameterOfBinaryTree

// 效率较低
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right), depth(root.Left)+depth(root.Right))
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(depth(root.Left), depth(root.Right))
}

func max(nums ...int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}
