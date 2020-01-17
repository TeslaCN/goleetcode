package p141_1

import "github.com/TeslaCN/goleetcode/util"

type ListNode = util.ListNode

func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	p1, p2 := head, head
	for p1 != nil && p2 != nil {
		p1 = p1.Next
		p2 = p2.Next
		if p2 != nil {
			p2 = p2.Next
		}
		if p1 == p2 {
			return true
		}
	}
	return false
}
