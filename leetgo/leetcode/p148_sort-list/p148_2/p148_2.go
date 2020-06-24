package p148_2

import "github.com/TeslaCN/goleetcode/util"

type ListNode = util.ListNode

var SortList = sortList

// Space: O(log(n))
// Runtime: O(n*log(n))
func sortList(head *ListNode) *ListNode {
	length := listLen(head)
	if length <= 1 {
		return head
	}

	// divide
	p := head
	for i := 0; i < length/2-1; i++ {
		p = p.Next
	}
	rightBank := p.Next
	p.Next = nil
	head = sortList(head)
	rightBank = sortList(rightBank)

	// conquer
	p = head
	var pre *ListNode
	for rightBank != nil {
		node := rightBank
		rightBank = rightBank.Next
		node.Next = nil
		for p != nil && p.Val < node.Val {
			pre = p
			p = p.Next
		}
		if pre == nil {
			node.Next = head
			head = node
			pre = node
		} else if p == nil {
			pre.Next = node
			node.Next = rightBank
			break
		} else {
			node.Next = pre.Next
			pre.Next = node
			pre = node
		}
	}
	return head
}

func listLen(head *ListNode) int {
	l := 0
	for head != nil {
		l++
		head = head.Next
	}
	return l
}
