package p142_1

import "github.com/TeslaCN/goleetcode/util"

type ListNode = util.ListNode

// Runtime: O(n)
// Space: O(n)
func DetectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	cycle := false
	p1, p2 := head, head
	for p1 != nil && p2 != nil {
		p1 = p1.Next
		p2 = p2.Next
		if p2 != nil {
			p2 = p2.Next
		}
		if p1 == p2 {
			cycle = true
			break
		}
	}
	if !cycle {
		return nil
	}

	set := make(map[*ListNode]bool)
	p := head
	for p != nil && !set[p] {
		set[p] = true
		p = p.Next
	}
	return p
}
