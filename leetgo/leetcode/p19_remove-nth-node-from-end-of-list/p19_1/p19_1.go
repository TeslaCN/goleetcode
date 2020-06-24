package p19_1

import "github.com/TeslaCN/goleetcode/util"

type ListNode = util.ListNode

var RemoveNthFromEnd = removeNthFromEnd

// Runtime: O(n)
// Space: O(n)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return head
	}
	list := make([]*ListNode, 0, 1<<4)
	p := head
	for p != nil {
		list = append(list, p)
		p = p.Next
	}
	l := len(list)
	if l == n {
		return head.Next
	}
	if n > l {
		// illegal: n
		return head
	}
	list[l-1-n].Next = list[l-n].Next
	return head
}
