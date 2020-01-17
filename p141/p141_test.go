package p141

import (
	"github.com/TeslaCN/goleetcode/p141/p141_1"
	"github.com/TeslaCN/goleetcode/util"
	"testing"
)

type ListNode = util.ListNode

type args struct {
	head *ListNode
}

var tests []struct {
	name string
	args args
	want bool
}

func init() {
	sample0 := util.CreateLinkedListInArray([]int{3, 2, 0, -4})
	sample1 := util.CreateLinkedListInArray([]int{1, 2})

	tests = append(tests, []struct {
		name string
		args args
		want bool
	}{
		{"sample-0", args{util.AppendLinkedList(sample0[3], sample0[1])}, true},
		{"sample-1", args{util.AppendLinkedList(sample1[1], sample1[0])}, true},
		{"sample-2", args{&ListNode{Val: 1}}, false},
	}...)
}

func TestHasCycle(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p141_1.HasCycle(tt.args.head); got != tt.want {
				t.Errorf("HasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
