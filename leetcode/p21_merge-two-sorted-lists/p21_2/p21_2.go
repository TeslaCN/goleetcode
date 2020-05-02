package p21_2

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
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

/*
执行用时 : 4 ms , 在所有 Go 提交中击败了 67.20% 的用户
内存消耗 : 2.6 MB , 在所有 Go 提交中击败了 63.64% 的用户
*/
