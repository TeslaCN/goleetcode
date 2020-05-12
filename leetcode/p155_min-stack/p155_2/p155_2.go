package p155_2

type ListNode struct {
	value int
	min   int
	next  *ListNode
}

type MinStack struct {
	head *ListNode
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	if this.head == nil {
		this.head = &ListNode{x, x, nil}
	} else {
		node := &ListNode{x, min(x, this.GetMin()), this.head}
		this.head = node
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (this *MinStack) Pop() {
	this.head = this.head.next
}

func (this *MinStack) Top() int {
	return this.head.value
}

func (this *MinStack) GetMin() int {
	return this.head.min
}

/*
执行用时 : 8 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 7.7 MB , 在所有 Go 提交中击败了 25.00% 的用户
*/
