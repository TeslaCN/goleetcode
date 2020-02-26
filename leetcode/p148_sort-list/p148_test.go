package p148_sort_list

import (
	"github.com/TeslaCN/goleetcode/leetcode/p148_sort-list/p148_1"
	"github.com/TeslaCN/goleetcode/leetcode/p148_sort-list/p148_2"
	"github.com/TeslaCN/goleetcode/util"
	"reflect"
	"testing"
)

type ListNode = util.ListNode

func Test_1(t *testing.T) {
	testSortList(t, p148_1.SortList)
}
func Test_2(t *testing.T) {
	testSortList(t, p148_2.SortList)
}

func testSortList(t *testing.T, f func(*ListNode) *ListNode) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"sample-2", args{util.CreateLinkedList([]int{3, 2, 1})}, util.CreateLinkedList([]int{1, 2, 3})},
		{"sample-0", args{util.CreateLinkedList([]int{4, 2, 1, 3})}, util.CreateLinkedList([]int{1, 2, 3, 4})},
		{"sample-1", args{util.CreateLinkedList([]int{-1, 5, 3, 4, 0})}, util.CreateLinkedList([]int{-1, 0, 3, 4, 5})},
		{"case-0", args{util.CreateLinkedList([]int{2, 1})}, util.CreateLinkedList([]int{1, 2})},
		{"case-1", args{util.CreateLinkedList([]int{1, 2})}, util.CreateLinkedList([]int{1, 2})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
