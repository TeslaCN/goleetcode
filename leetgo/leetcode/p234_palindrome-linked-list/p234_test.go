package p234_palindrome_linked_list

import (
	"github.com/TeslaCN/goleetcode/leetcode/p234_palindrome-linked-list/p234_1"
	"github.com/TeslaCN/goleetcode/util"
	"testing"
)

type ListNode = util.ListNode

type args struct {
	head *ListNode
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{"sample-0", args{util.CreateLinkedList([]int{1, 2})}, false},
	{"sample-1", args{util.CreateLinkedList([]int{1, 2, 2, 1})}, true},
	{"case-0", args{util.CreateLinkedList([]int{1, 2, 5, 2, 1})}, true},
	{"case-1", args{util.CreateLinkedList([]int{1, 6, 4, 7, 2, 2, 7, 4, 6, 1})}, true},
	{"case-2", args{util.CreateLinkedList([]int{1, 2, 3, 4, 5})}, false},
	{"case-3", args{util.CreateLinkedList([]int{1, 1})}, true},
	{"case-4", args{util.CreateLinkedList([]int{1, 2, 1})}, true},
}

func TestIsPalindrome(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p234_1.IsPalindrome(tt.args.head); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
