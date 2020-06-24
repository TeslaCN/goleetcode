package p23_1

import "github.com/TeslaCN/goleetcode/util"

type ListNode = util.ListNode

var MergeKLists = mergeKLists

// Space: O(1)
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	head := lists[0]

	for i := 1; i < len(lists); i++ {
		if head == nil {
			head = lists[i]
			continue
		}
		var cur, pre, p *ListNode = head, nil, lists[i]
		for cur != nil && p != nil {
			if cur.Val >= p.Val {
				if pre == nil {
					newHead := p
					p = p.Next
					newHead.Next = head
					head = newHead
					cur = newHead
				} else {
					node := p
					p = p.Next
					node.Next = cur
					pre.Next = node
					pre, cur = node, node.Next
					continue
				}
			}
			pre, cur = cur, cur.Next
		}
		if p != nil && pre != nil {
			pre.Next = p
		}
	}
	return head
}

/*
执行用时 : 116 ms , 在所有 Go 提交中击败了 31.42% 的用户
内存消耗 : 5.3 MB , 在所有 Go 提交中击败了 77.78% 的用户
*/
