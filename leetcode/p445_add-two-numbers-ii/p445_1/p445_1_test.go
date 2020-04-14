package p445_1

import (
	"github.com/TeslaCN/goleetcode/util"
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"sample-0", args{util.CreateLinkedList([]int{7, 2, 4, 3}), util.CreateLinkedList([]int{5, 6, 4})}, util.CreateLinkedList([]int{7, 8, 0, 7})},
		{"case-0", args{util.CreateLinkedList([]int{7, 2, 4, 3}), util.CreateLinkedList([]int{})}, util.CreateLinkedList([]int{7, 2, 4, 3})},
		{"case-1", args{util.CreateLinkedList([]int{9, 9, 9, 9, 9}), util.CreateLinkedList([]int{1})}, util.CreateLinkedList([]int{1, 0, 0, 0, 0, 0})},
		{"case-2", args{util.CreateLinkedList([]int{9, 9, 9, 9, 9}), util.CreateLinkedList([]int{9, 9, 9, 9, 9})}, util.CreateLinkedList([]int{1, 9, 9, 9, 9, 8})},
		{"case-3", args{util.CreateLinkedList([]int{9, 9, 9, 9, 9}), util.CreateLinkedList([]int{1, 0, 0, 0, 0, 1})}, util.CreateLinkedList([]int{2, 0, 0, 0, 0, 0})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
