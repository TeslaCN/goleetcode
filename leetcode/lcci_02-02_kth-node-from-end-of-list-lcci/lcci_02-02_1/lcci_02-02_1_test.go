package lcci_02_02_1

import (
	"github.com/TeslaCN/goleetcode/util"
	"testing"
)

func Test_kthToLast(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{util.CreateLinkedList([]int{1, 2, 3, 4, 5}), 2}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthToLast(tt.args.head, tt.args.k); got != tt.want {
				t.Errorf("kthToLast() = %v, want %v", got, tt.want)
			}
		})
	}
}
