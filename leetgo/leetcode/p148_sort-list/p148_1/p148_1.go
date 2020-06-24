package p148_1

import "github.com/TeslaCN/goleetcode/util"

type ListNode = util.ListNode

var SortList = sortList

// insert sort
// Space: O(1)
// Runtime: O(n^2)
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head
	head = head.Next
	newHead.Next = nil
	for head != nil {
		node := head
		head = head.Next
		node.Next = nil
		p := newHead
		var pre *ListNode

		for p != nil && node.Val > p.Val {
			pre = p
			p = p.Next
		}
		if pre == nil {
			node.Next = newHead
			newHead = node
		}
		if pre != nil {
			node.Next = pre.Next
			pre.Next = node
		}
	}
	return newHead
}
