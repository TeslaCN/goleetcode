package p617_1

import "github.com/TeslaCN/goleetcode/util"

type TreeNode = util.TreeNode

var SequenceTraversal = util.SequenceTraversal

// Deprecated
func MergeTrees(t1, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1Nodes := SequenceTraversal(t1)
	t2Nodes := SequenceTraversal(t2)
	if len(t2Nodes) > len(t1Nodes) {
		t1Nodes, t2Nodes = t2Nodes, t1Nodes
		t1, t2 = t2, t1
	}
	for i := 0; i < len(t2Nodes); i++ {
		if t2Nodes[i] != nil {
			if t1Nodes[i] == nil {
				parentIndex := (i - 1) / 2
				if i&1 == 1 {
					t1Nodes[parentIndex].Left = t2Nodes[i]
				} else {
					t1Nodes[parentIndex].Right = t2Nodes[i]
				}
			} else {
				t1Nodes[i].Val += t2Nodes[i].Val
			}
		}
	}
	return t1
}
