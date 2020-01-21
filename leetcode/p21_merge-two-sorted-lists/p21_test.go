package p21_merge_two_sorted_lists

import (
	"github.com/TeslaCN/goleetcode/leetcode/p21_merge-two-sorted-lists/p21_1"
	"github.com/TeslaCN/goleetcode/util"
	"reflect"
	"testing"
)

type ListNode = util.ListNode

type args struct {
	l1 *ListNode
	l2 *ListNode
}

var tests = []struct {
	name string
	args args
	want *ListNode
}{
	{"sample-0", args{util.CreateLinkedList([]int{1, 2, 4}), util.CreateLinkedList([]int{1, 3, 4})}, util.CreateLinkedList([]int{1, 1, 2, 3, 4, 4})},
	{"case-0", args{l1: util.CreateLinkedList([]int{1, 2, 3})}, util.CreateLinkedList([]int{1, 2, 3})},
	{"case-1", args{util.CreateLinkedList([]int{1}), util.CreateLinkedList([]int{1})}, util.CreateLinkedList([]int{1, 1})},
}

func TestMergeTwoLists(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p21_1.MergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
