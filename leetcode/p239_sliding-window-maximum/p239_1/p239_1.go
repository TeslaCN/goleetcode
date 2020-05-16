package p239_1

var MaxSlidingWindow = maxSlidingWindow

// 单调队列：
// 大于最大值直接清空原队列
// 小于最小值直接入列
// 其余情况小于目标值的出列
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 {
		return []int{}
	}
	maxValues, queue := make([]int, n-k+1), make([]int, 0, 1)
	for i := 0; i < k && i < n; i++ {
		queue = enqueue(queue, nums[i])
	}
	maxValues[0] = queue[0]
	for i := k; i < n; i++ {
		if len(queue) > 0 && queue[0] == nums[i-k] {
			queue = make([]int, 0, 1)
			for j := i - k + 1; j < i && j < n; j++ {
				queue = enqueue(queue, nums[j])
			}
		}
		queue = enqueue(queue, nums[i])
		maxValues[i-k+1] = queue[0]
	}
	return maxValues
}

func enqueue(queue []int, value int) []int {
	switch {
	case len(queue) == 0:
		queue = append(queue, value)
	case value > queue[0]:
		queue = []int{value}
	case value < queue[len(queue)-1]:
		queue = append(queue, value)
	default:
		cut := 0
		for ; cut < len(queue) && value < queue[cut]; cut++ {
		}
		queue[cut] = value
		queue = queue[:cut+1]
	}
	return queue
}

/*
执行用时 : 44 ms , 在所有 Go 提交中击败了 20.20% 的用户
内存消耗 : 7.2 MB , 在所有 Go 提交中击败了 25.00% 的用户
*/
