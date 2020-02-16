package p206_reverse_linked_list

import (
	"github.com/TeslaCN/goleetcode/leetcode/p206_reverse-linked-list/p206_1"
	"github.com/TeslaCN/goleetcode/leetcode/p206_reverse-linked-list/p206_2"
	"github.com/TeslaCN/goleetcode/util"
	"reflect"
	"testing"
)

type ListNode = util.ListNode

func Test_1(t *testing.T) {
	testReverseList(t, p206_1.ReverseList)
}

func Test_2(t *testing.T) {
	testReverseList(t, p206_2.ReverseList)
}

func testReverseList(t *testing.T, f func(*ListNode) *ListNode) {
	type args struct {
		head *ListNode
	}

	var tests = []struct {
		name string
		args args
		want *ListNode
	}{
		{"sample-0", args{util.CreateLinkedList([]int{1, 2, 3, 4, 5})}, util.CreateLinkedList([]int{5, 4, 3, 2, 1})},
		{"case-0", args{util.CreateLinkedList([]int{})}, util.CreateLinkedList([]int{})},
		{"case-1", args{util.CreateLinkedList([]int{1})}, util.CreateLinkedList([]int{1})},
		{"case-2", args{}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
