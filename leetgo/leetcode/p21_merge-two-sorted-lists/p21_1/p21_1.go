package p21_1

import "github.com/TeslaCN/goleetcode/util"

type ListNode = util.ListNode

var MergeTwoLists = mergeTwoLists

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l2.Val < l1.Val {
		l1, l2 = l2, l1
	}
	p1, p2 := l1, l2
	for p2 != nil {
		for p1.Next != nil && p1.Next.Val < p2.Val {
			p1 = p1.Next
		}
		p2Next := p2.Next
		p2.Next = p1.Next
		p1.Next = p2
		p2 = p2Next
	}
	return l1
}

/*
执行用时 : 4 ms , 在所有 Go 提交中击败了 67.20% 的用户
内存消耗 : 2.5 MB , 在所有 Go 提交中击败了 72.73% 的用户
*/
