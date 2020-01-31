package p437_1

import "github.com/TeslaCN/goleetcode/util"

type TreeNode = util.TreeNode

var PathSum = pathSum

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return find(root, 0, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func find(root *TreeNode, has int, target int) int {
	if root == nil {
		return 0
	}
	has += root.Val
	count := 0
	if has == target {
		count++
	}
	return count + find(root.Left, has, target) + find(root.Right, has, target)
}
