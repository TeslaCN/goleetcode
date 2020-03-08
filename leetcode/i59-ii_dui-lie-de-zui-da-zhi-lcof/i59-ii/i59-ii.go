package i59_ii

type MaxQueue struct {
	values   []int
	monotone []int
	start    int
}

func Constructor() MaxQueue {
	return MaxQueue{start: 0}
}

func (this *MaxQueue) Max_value() int {
	if this.start >= len(this.values) {
		return -1
	}
	return this.values[this.monotone[0]]
}

func (this *MaxQueue) Push_back(value int) {
	this.values = append(this.values, value)
	index := len(this.values) - 1
	if len(this.monotone) == 0 {
		this.monotone = append(this.monotone, index)
	} else {
		if value >= this.values[this.monotone[0]] {
			this.monotone = []int{index}
		} else {
			i := len(this.monotone) - 1
			for ; i > 0 && this.values[index] >= this.values[this.monotone[i]]; i-- {
			}
			this.monotone = append(this.monotone[:i+1], index)
		}
	}
}

func (this *MaxQueue) Pop_front() int {
	if this.start >= len(this.values) {
		return -1
	}
	pop := this.values[this.start]
	this.start++
	if this.start > this.monotone[0] {
		this.monotone = this.monotone[1:]
	}
	return pop
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
