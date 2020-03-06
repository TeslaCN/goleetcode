package p225_1

import "container/list"

type MyStack struct {
	list *list.List
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{list: list.New()}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.list.PushBack(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	for i := 0; i < this.list.Len()-1; i++ {
		this.list.PushBack(this.list.Remove(this.list.Front()))
	}
	return this.list.Remove(this.list.Front()).(int)
}

/** Get the top element. */
func (this *MyStack) Top() int {
	for i := 0; i < this.list.Len()-1; i++ {
		this.list.PushBack(this.list.Remove(this.list.Front()))
	}
	top := this.list.Front().Value.(int)
	this.list.PushBack(this.list.Remove(this.list.Front()))
	return top
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.list.Len() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
