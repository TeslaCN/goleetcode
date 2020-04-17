package p23_merge_k_sorted_lists

import (
	"github.com/TeslaCN/goleetcode/leetcode/p23_merge-k-sorted-lists/p23_1"
	"github.com/TeslaCN/goleetcode/leetcode/p23_merge-k-sorted-lists/p23_2"
	"github.com/TeslaCN/goleetcode/util"
	"reflect"
	"testing"
)

type ListNode = util.ListNode

func Test_1(t *testing.T) {
	testMergeKLists(t, p23_1.MergeKLists)
}

func Test_2(t *testing.T) {
	testMergeKLists(t, p23_2.MergeKLists)
}

func testMergeKLists(t *testing.T, f func([]*ListNode) *ListNode) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"sample-0", args{[]*ListNode{
			util.CreateLinkedList([]int{1, 4, 5}),
			util.CreateLinkedList([]int{1, 3, 4}),
			util.CreateLinkedList([]int{2, 6}),
		}}, util.CreateLinkedList([]int{1, 1, 2, 3, 4, 4, 5, 6})},
		{"sample-1", args{[]*ListNode{
			util.CreateLinkedList([]int{}),
			util.CreateLinkedList([]int{}),
			util.CreateLinkedList([]int{}),
		}}, util.CreateLinkedList([]int{})},
		{"case-0", args{[]*ListNode{
			util.CreateLinkedList([]int{9}),
			util.CreateLinkedList([]int{1, 3, 4}),
			util.CreateLinkedList([]int{2, 6}),
		}}, util.CreateLinkedList([]int{1, 2, 3, 4, 6, 9})},
		{"case-1", args{[]*ListNode{
			util.CreateLinkedList([]int{1}),
			util.CreateLinkedList([]int{2}),
			util.CreateLinkedList([]int{3}),
			util.CreateLinkedList([]int{2}),
			util.CreateLinkedList([]int{1}),
		}}, util.CreateLinkedList([]int{1, 1, 2, 2, 3})},
		{"case-2", args{[]*ListNode{
			util.CreateLinkedList([]int{1}),
			util.CreateLinkedList([]int{}),
			util.CreateLinkedList([]int{3}),
			util.CreateLinkedList([]int{2}),
			util.CreateLinkedList([]int{1}),
		}}, util.CreateLinkedList([]int{1, 1, 2, 3})},
		{"case-3", args{[]*ListNode{
			util.CreateLinkedList([]int{}),
			util.CreateLinkedList([]int{}),
			util.CreateLinkedList([]int{2}),
		}}, util.CreateLinkedList([]int{2})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
