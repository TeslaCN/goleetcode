package p206_1

import "github.com/TeslaCN/goleetcode/util"

type ListNode = util.ListNode

// 递归实现
func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	return reverse(nil, head)
}

func reverse(pre, next *ListNode) *ListNode {
	if next == nil {
		return pre
	}
	newHead := reverse(next, next.Next)
	next.Next = pre
	return newHead
}
