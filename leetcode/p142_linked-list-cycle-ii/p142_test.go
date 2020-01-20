package p142_linked_list_cycle_ii

import (
	"github.com/TeslaCN/goleetcode/leetcode/p142_linked-list-cycle-ii/p142_1"
	"github.com/TeslaCN/goleetcode/leetcode/p142_linked-list-cycle-ii/p142_2"
	"github.com/TeslaCN/goleetcode/leetcode/p142_linked-list-cycle-ii/p142_3"
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

	case0 := util.CreateLinkedListInArray([]int{-1, 0, 1, 2, 3})
	util.AppendLinkedList(case0[len(case0)-1], case0[1])

	case1 := util.CreateLinkedListInArray([]int{-1, 0, 1, 2, 3, 4})
	util.AppendLinkedList(case1[len(case1)-1], case1[1])

	case2 := &ListNode{Val: 0}
	case2.Next = case2

	sample3 := util.CreateLinkedListInArray([]int{-1, -7, 7, -4, 19, 6, -9, -5, -2, -5})
	sample3[len(sample3)-1].Next = sample3[len(sample3)-1]

	tests = append(tests, []struct {
		name string
		args args
		want *ListNode
	}{
		{"sample-0", args{sample0[0]}, sample0[1]},
		{"sample-1", args{sample1[0]}, sample1[0]},
		{"sample-2", args{&ListNode{Val: 1}}, nil},
		{"sample-3", args{sample3[0]}, sample3[len(sample3)-1]},
		{"case-0", args{case0[0]}, case0[1]},
		{"case-1", args{case1[0]}, case1[1]},
		{"case-2", args{case2}, case2},
	}...)
}

func Test_1(t *testing.T) {
	testDetectCycle(t, p142_1.DetectCycle)
}

func Test_2(t *testing.T) {
	testDetectCycle(t, p142_2.DetectCycle)
}

func Test_3(t *testing.T) {
	testDetectCycle(t, p142_3.DetectCycle)
}

func testDetectCycle(t *testing.T, f func(*ListNode) *ListNode) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
