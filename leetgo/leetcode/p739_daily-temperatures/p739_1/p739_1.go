package p739_1

var DailyTemperatures = dailyTemperatures

// Brute Force
// Runtime: O(n^2)
// Space: O(1)
func dailyTemperatures(T []int) []int {
	n := len(T)
	for i := 0; i < n; i++ {
		changed := false
		for j := i + 1; j < n; j++ {
			if T[j] > T[i] {
				T[i], changed = j-i, true
				break
			}
		}
		if !changed {
			T[i] = 0
		}
	}
	return T
}

/*
执行用时 : 644 ms , 在所有 Go 提交中击败了 15.82% 的用户
内存消耗 : 6.4 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
