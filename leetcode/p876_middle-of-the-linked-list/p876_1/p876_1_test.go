package p876_1

import (
	"github.com/TeslaCN/goleetcode/util"
	"reflect"
	"testing"
)

func Test_middleNode(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"sample-0", args{util.CreateLinkedList([]int{1, 2, 3, 4, 5})}, util.CreateLinkedList([]int{3, 4, 5})},
		{"sample-1", args{util.CreateLinkedList([]int{1, 2, 3, 4, 5, 6})}, util.CreateLinkedList([]int{4, 5, 6})},
		{"case-0", args{util.CreateLinkedList([]int{})}, util.CreateLinkedList([]int{})},
		{"case-1", args{util.CreateLinkedList([]int{1})}, util.CreateLinkedList([]int{1})},
		{"case-2", args{util.CreateLinkedList([]int{1, 2})}, util.CreateLinkedList([]int{2})},
		{"case-3", args{util.CreateLinkedList([]int{1, 2, 3})}, util.CreateLinkedList([]int{2, 3})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
