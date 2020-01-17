package p142_3

import (
	"github.com/TeslaCN/goleetcode/util"
	"reflect"
)

type ListNode = util.ListNode

// Runtime: O(n)
// Space: O(1)
// 利用地址判断，只适用于顺序分配内存的链表
func DetectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	p := head
	for p != nil {
		if p.Next == nil {
			return nil
		}
		if reflect.ValueOf(p.Next).Pointer() <= reflect.ValueOf(p).Pointer() {
			return p.Next
		}
		p = p.Next
	}
	return nil
}
