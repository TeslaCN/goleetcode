package p160_1

import (
	"github.com/TeslaCN/goleetcode/util"
)

type ListNode = util.ListNode

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	lengthA := lengthOfLinkedList(headA)
	lengthB := lengthOfLinkedList(headB)

	if lengthA < lengthB {
		headA, headB = headB, headA
		lengthA, lengthB = lengthB, lengthA
	}
	pA := headA
	pB := headB

	for i := 0; i < lengthA-lengthB; i++ {
		pA = pA.Next
	}

	for pA != nil {
		if pA == pB {
			return pA
		}
		pA = pA.Next
		pB = pB.Next
	}
	return nil
}

func lengthOfLinkedList(head *ListNode) int {
	length := 0
	if head == nil {
		return length
	}
	for ; head != nil; head = head.Next {
		length++
	}
	return length
}
