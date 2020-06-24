package lcci_02_02_1

import "github.com/TeslaCN/goleetcode/util"

type ListNode = util.ListNode

// k is valid
func kthToLast(head *ListNode, k int) int {
	pre, after := head, head
	for i := 0; i < k && pre != nil; i++ {
		pre = pre.Next
	}
	for pre != nil {
		pre, after = pre.Next, after.Next
	}
	return after.Val
}

/*
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 2.2 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
