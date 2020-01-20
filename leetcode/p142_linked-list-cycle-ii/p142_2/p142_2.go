package p142_2

import (
	"github.com/TeslaCN/goleetcode/util"
)

type ListNode = util.ListNode

// Runtime: O(n^2)
// Space: O(1)
func DetectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	p1, p2 := head, head
	preLen := 0
	cycle := false
	cycleLen := 0
	for p1 != nil && p2 != nil {
		p1 = p1.Next
		if p2 = p2.Next; p2 != nil {
			p2 = p2.Next
		}
		if p1 == p2 {
			if !cycle {
				cycle = true
			} else {
				break
			}
		}
		if cycle {
			cycleLen++
		} else {
			preLen++
		}
	}
	//log.Printf("pre len: %v", preLen)
	if !cycle {
		return nil
	}
	//log.Printf("cycle len: %v", cycleLen)
	maxLen := preLen + cycleLen + 1
	p := head
	for i := 0; i < maxLen; i++ {
		node := p
		for j := 0; j < maxLen; j++ {
			if node.Next == p {
				return p
			}
			node = node.Next
		}
		p = p.Next
	}
	return nil
}
