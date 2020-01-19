package p160

import (
	"github.com/TeslaCN/goleetcode/leetcode/p160/p160_1"
	"github.com/TeslaCN/goleetcode/util"
	"reflect"
	"testing"
)

type ListNode = util.ListNode

type args struct {
	headA *p160_1.ListNode
	headB *p160_1.ListNode
}

var tests []struct {
	name string
	args args
	want *p160_1.ListNode
}

func init() {
	case0 := util.CreateLinkedList([]int{8, 4, 5})
	tests = append(tests, []struct {
		name string
		args args
		want *p160_1.ListNode
	}{
		{
			"case-0",
			args{
				util.AppendLinkedList(util.CreateLinkedList([]int{4, 1}), case0),
				util.AppendLinkedList(util.CreateLinkedList([]int{5, 0, 1}), case0),
			},
			util.CreateLinkedList([]int{8, 4, 5}),
		},
	}...)
}

func Test_1(t *testing.T) {
	testGetIntersectionNode(t, p160_1.GetIntersectionNode)
}

func testGetIntersectionNode(t *testing.T, f func(*ListNode, *ListNode) *ListNode) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := f(tt.args.headA, tt.args.headB)
			t.Logf("%v %v => %v", tt.args.headA, tt.args.headB, got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
