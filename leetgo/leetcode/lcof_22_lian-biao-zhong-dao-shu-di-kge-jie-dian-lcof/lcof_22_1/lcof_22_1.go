package lcof_22_1

import "github.com/TeslaCN/goleetcode/util"

type ListNode = util.ListNode

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || k < 1 {
		return nil
	}
	pre, after := head, head
	for i := 0; i < k; i++ {
		if pre != nil {
			pre = pre.Next
		} else {
			return nil
		}
	}
	for pre != nil {
		pre, after = pre.Next, after.Next
	}
	return after
}

/*
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 2.2 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
