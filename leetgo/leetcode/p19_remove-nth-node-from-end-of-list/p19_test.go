package p19_remove_nth_node_from_end_of_list

import (
	"github.com/TeslaCN/goleetcode/leetcode/p19_remove-nth-node-from-end-of-list/p19_1"
	"github.com/TeslaCN/goleetcode/leetcode/p19_remove-nth-node-from-end-of-list/p19_2"
	"github.com/TeslaCN/goleetcode/util"
	"reflect"
	"testing"
)

type ListNode = util.ListNode

func Test_1(t *testing.T) {
	testRemoveNthFromEnd(t, p19_1.RemoveNthFromEnd)
}

func Test_2(t *testing.T) {
	testRemoveNthFromEnd(t, p19_2.RemoveNthFromEnd)
}

func testRemoveNthFromEnd(t *testing.T, f func(*ListNode, int) *ListNode) {
	type args struct {
		head *ListNode
		n    int
	}

	var tests = []struct {
		name string
		args args
		want *ListNode
	}{
		{"sample-0", args{util.CreateLinkedList([]int{1, 2, 3, 4, 5}), 2}, util.CreateLinkedList([]int{1, 2, 3, 5})},
		{"sample-1", args{util.CreateLinkedList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), 7}, util.CreateLinkedList([]int{1, 2, 3, 5, 6, 7, 8, 9, 10})},
		{"case-0", args{util.CreateLinkedList([]int{1, 2, 3, 4, 5}), 1}, util.CreateLinkedList([]int{1, 2, 3, 4})},
		{"case-1", args{util.CreateLinkedList([]int{1, 2, 3, 4, 5}), 5}, util.CreateLinkedList([]int{2, 3, 4, 5})},
		{"case-2", args{util.CreateLinkedList([]int{1}), 1}, util.CreateLinkedList([]int{})},
		{"case-3", args{util.CreateLinkedList([]int{1, 2}), 1}, util.CreateLinkedList([]int{1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
