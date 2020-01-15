package util

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLinkedList(t *testing.T) {
	ints := []int{0, 1, 2, 3, 4, 3, 2, 1, 0, -1}
	head := CreateLinkedList(ints)
	PrintLinkedList(head)
	headB := CreateLinkedList(ints)

	fmt.Printf("DeepEqual: %v %v %v\n", head, headB, reflect.DeepEqual(head, headB))
	fmt.Printf("ValueEqual: %v  %v %v\n", head, headB, ValueEquals(head, headB))
}

func TestValueEquals(t *testing.T) {
	nodeA := &ListNode{233, nil}
	nodeB := &ListNode{233, nil}
	t.Logf("%v %v =? %v", nodeA, nodeB, nodeA == nodeB)
	t.Logf("%v %v deepEqual? %v", nodeA, nodeB, reflect.DeepEqual(nodeA, nodeB))
}
