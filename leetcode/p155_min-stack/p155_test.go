package p155_min_stack

import (
	"fmt"
	"github.com/TeslaCN/goleetcode/leetcode/p155_min-stack/p155_1"
	"testing"
)

func TestArray(t *testing.T) {
	a := make([]int, 0)
	for i := 0; i < 128; i++ {
		fmt.Printf("%v -> %v\n", i, cap(a))
		a = append(a, i)
	}
}

func TestMinStack(t *testing.T) {
	minStack := p155_1.Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
}
