package i57_ii_1

// 滑动窗口
func findContinuousSequence(target int) [][]int {
	sets := make([][]int, 0)
	left, right := 1, 2
	sum := 3
	for right <= (target+1)>>1 {
		if sum < target {
			right++
			sum += right
		} else if sum > target {
			sum -= left
			left++
		} else {
			set := make([]int, 0)
			for i := left; i <= right; i++ {
				set = append(set, i)
			}
			sets = append(sets, set)
			right++
			sum += right
		}
	}
	return sets
}
