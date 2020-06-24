package p739_2

var DailyTemperatures = dailyTemperatures

func dailyTemperatures(T []int) []int {
	n := len(T)
	result := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j += result[j] {
			if T[i] < T[j] {
				result[i] = j - i
				break
			}
			if result[j] == 0 {
				result[i] = 0
				break
			}
		}
	}
	return result
}

/*
执行用时 : 60 ms , 在所有 Go 提交中击败了 90.51% 的用户
内存消耗 : 6.9 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
