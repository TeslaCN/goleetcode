package util

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateLinkedList(values []int) *ListNode {
	if values == nil || len(values) < 1 {
		return nil
	}
	headNode := &ListNode{values[0], nil}
	p := headNode
	for i := 1; i < len(values); i++ {
		p.Next = &ListNode{values[i], nil}
		p = p.Next
	}
	return headNode
}

func AppendLinkedList(head, appendixHead *ListNode) *ListNode {
	if head == nil || appendixHead == nil {
		return head
	}
	p := head
	for p.Next != nil {
		p = p.Next
	}
	p.Next = appendixHead
	return head
}

func PrintLinkedList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func ValueEquals(headA, headB *ListNode) bool {
	if headA == headB {
		return true
	}
	if headA == nil || headB == nil {
		return false
	}
	pA := headA
	pB := headB
	for pA != nil && pB != nil {
		if pA.Val != pB.Val {
			return false
		}
		pA = pA.Next
		pB = pB.Next
	}
	return true
}
