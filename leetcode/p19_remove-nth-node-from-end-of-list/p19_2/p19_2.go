package p19_2

import "github.com/TeslaCN/goleetcode/util"

type ListNode = util.ListNode

var RemoveNthFromEnd = removeNthFromEnd

// Runtime: O(n)
// Space: O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return head
	}
	p1, p2 := head, head
	for i := 0; i < n; i++ {
		if p2 == nil {
			// illegal: n
			return head
		}
		p2 = p2.Next
	}
	if p2 == nil {
		return head.Next
	}

	for p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	p1.Next = p1.Next.Next
	return head
}
