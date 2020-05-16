package p25_1

import "github.com/TeslaCN/goleetcode/util"

type ListNode = util.ListNode

// Runtime: O(n)
// Space: O(1)
func reverseKGroup(head *ListNode, k int) *ListNode {
	n := listLen(head)
	if n <= 1 || k < 2 {
		return head
	}
	grouped := n / k * k
	var fixedNode, pre, cur, next *ListNode = nil, nil, head, head.Next
	if grouped != n {
		fixedNode = head
		for i := 0; i < grouped; i++ {
			fixedNode = fixedNode.Next
		}
	}
	for i := 0; i < grouped && cur != nil; i++ {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	head = pre
	pre, cur, next = fixedNode, head, nil
	for i := 0; i < n/k; i++ {
		curPointer := cur
		for j := 0; j < k-1; j++ {
			curPointer = curPointer.Next
		}
		next = curPointer.Next
		curPointer.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func listLen(root *ListNode) int {
	l := 0
	for ; root != nil; root = root.Next {
		l++
	}
	return l
}

/*
执行用时 : 4 ms , 在所有 Go 提交中击败了 93.25% 的用户
内存消耗 : 3.6 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
