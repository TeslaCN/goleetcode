package p225_1

import (
	"testing"
)

func TestMyStack(t *testing.T) {
	stack := Constructor()
	t.Log(stack.Empty() == true)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	t.Log(stack.Empty() == false)
	t.Log(stack.Top() == 3)
	t.Log(stack.Pop() == 3)
	t.Log(stack.Top() == 2)
	stack.Pop()
	stack.Pop()
	//stack.Pop() // invalid
}
