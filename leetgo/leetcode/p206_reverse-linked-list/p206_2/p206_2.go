package p206_2

import "github.com/TeslaCN/goleetcode/util"

type ListNode = util.ListNode

// 迭代实现
func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, mid *ListNode = nil, head
	for mid != nil {
		next := mid.Next
		mid.Next = pre
		pre = mid
		mid = next
	}
	return pre
}
