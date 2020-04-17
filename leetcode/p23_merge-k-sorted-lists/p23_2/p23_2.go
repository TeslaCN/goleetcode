package p23_2

import "github.com/TeslaCN/goleetcode/util"

type ListNode = util.ListNode

var MergeKLists = mergeKLists

// Space: O(1)
func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	// Remove nil
	from, to := 0, 0
	for from < len(lists) {
		for from < len(lists) && lists[from] == nil {
			from++
		}
		if from >= len(lists) {
			break
		}
		lists[to] = lists[from]
		to++
		from++
	}
	lists = lists[:to]

	// Build heap
	for i := (len(lists) - 1) >> 1; i >= 0; i-- {
		adjustHeap(lists, i, len(lists))
	}
	var newHead, p *ListNode
	end := len(lists)
	for end > 0 {
		node := lists[0]
		lists[0] = lists[0].Next
		if lists[0] == nil {
			lists[0] = lists[end-1]
			end--
		}
		if newHead == nil {
			newHead = node
			node.Next = nil
			p = newHead
		} else if p != nil {
			p.Next = node
			p = p.Next
			p.Next = nil
		}
		adjustHeap(lists, 0, end)
	}
	return newHead
}

func adjustHeap(lists []*ListNode, target, end int) {
	node := lists[target]
	for child := target<<1 + 1; child < end; child = child<<1 + 1 {
		if child+1 < end && lists[child+1].Val < lists[child].Val {
			child++
		}
		if lists[child].Val < node.Val {
			lists[target], target = lists[child], child
		}
	}
	lists[target] = node
}

/*
执行用时 : 4 ms , 在所有 Go 提交中击败了 99.70% 的用户
内存消耗 : 5.3 MB , 在所有 Go 提交中击败了 100.00% 的用户

执行用时 : 12 ms , 在所有 Go 提交中击败了 77.78% 的用户
内存消耗 : 5.3 MB , 在所有 Go 提交中击败了 77.78% 的用户
*/
