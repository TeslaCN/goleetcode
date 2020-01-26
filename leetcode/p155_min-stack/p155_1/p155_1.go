package p155_1

type MinStack struct {
	a      []int
	min    []int
	cur    int
	minCur int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		a:      make([]int, 1<<6),
		min:    make([]int, 1<<6),
		cur:    0,
		minCur: 0,
	}
}

func (s *MinStack) Push(x int) {
	if s.cur == cap(s.a) {
		s.a = append(s.a, 0)
	}
	if s.minCur == cap(s.min) {
		s.min = append(s.min, 0)
	}
	s.a[s.cur] = x
	s.cur++
	if s.minCur == 0 {
		s.min[s.minCur] = x
	} else if x > s.min[s.minCur-1] {
		s.min[s.minCur] = s.min[s.minCur-1]
	} else {
		s.min[s.minCur] = x
	}
	s.minCur++
}

func (s *MinStack) Pop() {
	s.cur--
	s.minCur--
}

func (s *MinStack) Top() int {
	if s.cur > 0 {
		return s.a[s.cur-1]
	}
	return 0
}

func (s *MinStack) GetMin() int {
	if s.minCur > 0 {
		return s.min[s.minCur-1]
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
