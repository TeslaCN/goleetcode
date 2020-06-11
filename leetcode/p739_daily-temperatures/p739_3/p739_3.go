package p739_3

var DailyTemperatures = dailyTemperatures

func dailyTemperatures(T []int) []int {
	n := len(T)
	if n == 0 {
		return []int{}
	}
	ans, stack := make([]int, n), []int{0}
	for i := 1; i < n; i++ {
		for len(stack) > 0 && T[i] > T[stack[len(stack)-1]] {
			pop := stack[len(stack)-1]
			stack, ans[pop] = stack[:len(stack)-1], i-pop
		}
		stack = append(stack, i)
	}
	for ; len(stack) > 0; stack = stack[:len(stack)-1] {
		ans[stack[len(stack)-1]] = 0
	}
	return ans
}

/*
执行用时 : 60 ms , 在所有 Go 提交中击败了 90.32% 的用户
内存消耗 : 7.1 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
