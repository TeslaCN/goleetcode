package p977_2

var SortedSquares = sortedSquares

// Runtime: O(n)
// Space: O(n)
func sortedSquares(A []int) []int {
	n := len(A)
	left, right, ans, cur := 0, n-1, make([]int, n), n-1
	for left <= right {
		leftValue, rightValue := A[left], A[right]
		if leftValue < 0 {
			leftValue = -leftValue
		}
		if rightValue < 0 {
			rightValue = -rightValue
		}
		if leftValue >= rightValue {
			ans[cur] = leftValue * leftValue
			left++
		} else {
			ans[cur] = rightValue * rightValue
			right--
		}
		cur--
	}
	return ans
}

/*
执行用时 : 32 ms , 在所有 Go 提交中击败了 93.93% 的用户
内存消耗 : 6.5 MB , 在所有 Go 提交中击败了 25.00% 的用户
*/
