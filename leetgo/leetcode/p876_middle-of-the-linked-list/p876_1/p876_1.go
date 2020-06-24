package p876_1

import "github.com/TeslaCN/goleetcode/util"

type ListNode = util.ListNode

func middleNode(head *ListNode) *ListNode {
	count := 0
	for p := head; p != nil; p = p.Next {
		count++
	}
	for i := 0; i < count>>1; i++ {
		head = head.Next
	}
	return head
}
