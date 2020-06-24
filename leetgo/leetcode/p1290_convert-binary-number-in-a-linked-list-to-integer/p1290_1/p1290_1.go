package p1290_1

import "github.com/TeslaCN/goleetcode/util"

type ListNode = util.ListNode

func getDecimalValue(head *ListNode) int {
	num := 0
	for ; head != nil; head = head.Next {
		num = num<<1 | head.Val
	}
	return num
}
