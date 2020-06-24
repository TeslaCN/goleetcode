package p445_1

import "github.com/TeslaCN/goleetcode/util"

type ListNode = util.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	nodes1, nodes2 := make([]*ListNode, 0), make([]*ListNode, 0)
	for p := l1; p != nil; p = p.Next {
		nodes1 = append(nodes1, p)
	}
	for p := l2; p != nil; p = p.Next {
		nodes2 = append(nodes2, p)
	}
	i, j := len(nodes1)-1, len(nodes2)-1
	if i < j {
		i, j = j, i
		nodes1, nodes2 = nodes2, nodes1
	}

	var head *ListNode
	sum := 0
	for i >= 0 && j >= 0 {
		sum = nodes1[i].Val + nodes2[j].Val + sum/10
		head = &ListNode{
			Val:  sum % 10,
			Next: head,
		}
		i--
		j--
	}
	for ; i >= 0; i-- {
		sum = nodes1[i].Val + sum/10
		head = &ListNode{
			Val:  sum % 10,
			Next: head,
		}
	}
	if sum >= 10 {
		head = &ListNode{
			Val:  1,
			Next: head,
		}
	}
	return head
}

/*
执行用时 : 12 ms , 在所有 Go 提交中击败了 80.38% 的用户
内存消耗 : 5.7 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
