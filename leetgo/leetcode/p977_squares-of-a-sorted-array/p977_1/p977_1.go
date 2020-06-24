package p977_1

var SortedSquares = sortedSquares

// Runtime: O(n)
// Space: O(n)
func sortedSquares(A []int) []int {
	n := len(A)
	left := -1
	for ; left < n-1 && A[left+1] < 0; left++ {
	}
	right, ans, cur := left+1, make([]int, n), 0
	for left >= 0 && right < n {
		if -A[left] <= A[right] {
			ans[cur] = A[left] * A[left]
			left--
		} else {
			ans[cur] = A[right] * A[right]
			right++
		}
		cur++
	}
	for left >= 0 {
		ans[cur] = A[left] * A[left]
		cur++
		left--
	}
	for right < n {
		ans[cur] = A[right] * A[right]
		cur++
		right++
	}
	return ans
}

/*
执行用时 : 36 ms , 在所有 Go 提交中击败了 78.50% 的用户
内存消耗 : 6.5 MB , 在所有 Go 提交中击败了 25.00% 的用户
*/
