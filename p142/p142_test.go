package p142

import (
	"github.com/TeslaCN/goleetcode/p142/p142_1"
	"github.com/TeslaCN/goleetcode/util"
	"reflect"
	"testing"
)

type ListNode = util.ListNode

type args struct {
	head *ListNode
}

var tests []struct {
	name string
	args args
	want *ListNode
}

func init() {
	sample0 := util.CreateLinkedListInArray([]int{3, 2, 0, -4})
	util.AppendLinkedList(sample0[3], sample0[1])

	sample1 := util.CreateLinkedListInArray([]int{1, 2})
	util.AppendLinkedList(sample1[1], sample1[0])
	tests = append(tests, []struct {
		name string
		args args
		want *ListNode
	}{
		{"sample-0", args{sample0[0]}, sample0[1]},
		{"sample-1", args{sample1[0]}, sample1[0]},
		{"sample-2", args{&ListNode{Val: 1}}, nil},
	}...)
}

func TestDetectCycle(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p142_1.DetectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
