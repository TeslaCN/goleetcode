package p234_1

import "github.com/TeslaCN/goleetcode/util"

type ListNode = util.ListNode

// 逆转链表前半部分
// Runtime: O(n)
// Space: O(1)
func IsPalindrome(head *ListNode) bool {
	length := lenOf(head)
	if length <= 1 {
		return true
	}
	var pre, cur, next *ListNode = nil, nil, head
	for i := 0; i < length/2; i++ {
		cur = next
		next = cur.Next
		cur.Next = pre
		pre = cur
	}

	head1, head2 := cur, next
	if length&1 == 1 {
		head2 = head2.Next
	}
	p1, p2 := head1, head2
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

func lenOf(head *ListNode) int {
	p, length := head, 0
	for ; p != nil; p = p.Next {
		length++
	}
	return length
}
