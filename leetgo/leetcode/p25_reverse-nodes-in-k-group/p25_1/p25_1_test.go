package p25_1

import (
	"github.com/TeslaCN/goleetcode/util"
	"reflect"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"sample-0", args{util.CreateLinkedList([]int{1, 2, 3, 4, 5}), 2}, util.CreateLinkedList([]int{2, 1, 4, 3, 5})},
		{"case-0", args{util.CreateLinkedList([]int{1, 2, 3, 4, 5, 6, 7}), 2}, util.CreateLinkedList([]int{2, 1, 4, 3, 6, 5, 7})},
		{"case-1", args{util.CreateLinkedList([]int{1, 2, 3, 4, 5, 6, 7}), 3}, util.CreateLinkedList([]int{3, 2, 1, 6, 5, 4, 7})},
		{"sample-1", args{util.CreateLinkedList([]int{1, 2, 3, 4, 5}), 3}, util.CreateLinkedList([]int{3, 2, 1, 4, 5})},
		{"case-2", args{util.CreateLinkedList([]int{1, 2, 3, 4, 5, 6, 7}), 1}, util.CreateLinkedList([]int{1, 2, 3, 4, 5, 6, 7})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
