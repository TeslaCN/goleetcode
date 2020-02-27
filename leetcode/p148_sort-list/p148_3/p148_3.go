package p148_3

import (
	"github.com/TeslaCN/goleetcode/util"
)

type ListNode = util.ListNode

var SortList = sortList

// Space: O(1)
// Runtime: O(n*log(n))
func sortList(head *ListNode) *ListNode {
	length := listLen(head)
	if length <= 1 {
		return head
	}

	for step := 2; ; step <<= 1 {
		var pre, next *ListNode = nil, head
		for next != nil {
			var left, right *ListNode = next, nil
			p := next
			for i := 0; p != nil && i < step/2-1; i++ {
				p = p.Next
			}
			if p != nil {
				right = p.Next
				p.Next = nil
			}

			p = right
			for i := 0; p != nil && i < step/2-1; i++ {
				p = p.Next
			}
			if p != nil {
				next = p.Next
				p.Next = nil
			} else {
				next = nil
			}

			newHead := merge(left, right)
			if pre == nil {
				head = newHead
			} else {
				pre.Next = newHead
			}
			p = newHead
			for p != nil {
				pre = p
				p = p.Next
			}
		}
		if step >= length {
			break
		}
	}
	return head
}

func merge(left, right *ListNode) *ListNode {
	if right == nil {
		return left
	}
	if left == nil {
		return right
	}

	var p, pre *ListNode = left, nil
	for right != nil {
		node := right
		right = right.Next
		node.Next = nil
		for p != nil && node.Val > p.Val {
			pre = p
			p = p.Next
		}
		if pre == nil {
			node.Next = left
			left = node
			pre = node
		} else if p == nil {
			pre.Next = node
			pre = node
		} else {
			node.Next = pre.Next
			pre.Next = node
			pre = node
		}
	}
	return left
}

func listLen(head *ListNode) int {
	count := 0
	for head != nil {
		count++
		head = head.Next
	}
	return count
}
