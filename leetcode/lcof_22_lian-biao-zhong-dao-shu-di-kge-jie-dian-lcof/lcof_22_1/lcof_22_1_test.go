package lcof_22_1

import (
	"github.com/TeslaCN/goleetcode/util"
	"testing"
)

func Test_getKthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"sample-0", args{util.CreateLinkedList([]int{1, 2, 3, 4, 5, 6}), 3}, &ListNode{Val: 4}},
		{"case-0", args{util.CreateLinkedList([]int{1, 2, 3, 4, 5, 6}), 6}, &ListNode{Val: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKthFromEnd(tt.args.head, tt.args.k); got.Val != tt.want.Val {
				t.Errorf("getKthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
