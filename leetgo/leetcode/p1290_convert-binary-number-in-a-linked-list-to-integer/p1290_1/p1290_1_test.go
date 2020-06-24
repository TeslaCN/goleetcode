package p1290_1

import (
	"github.com/TeslaCN/goleetcode/util"
	"testing"
)

func Test_getDecimalValue(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{util.CreateLinkedList([]int{1, 0, 1})}, 5},
		{"sample-1", args{util.CreateLinkedList([]int{0})}, 0},
		{"sample-2", args{util.CreateLinkedList([]int{1})}, 1},
		{"sample-3", args{util.CreateLinkedList([]int{0, 0})}, 0},
		{"sample-4", args{util.CreateLinkedList([]int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0})}, 18880},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDecimalValue(tt.args.head); got != tt.want {
				t.Errorf("getDecimalValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
