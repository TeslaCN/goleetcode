package p977_3

import "sort"

var SortedSquares = sortedSquares

// Runtime: O(nlog(n))
// Space: O(1) ? O(n)
func sortedSquares(A []int) []int {
	n := len(A)
	for i := 0; i < n; i++ {
		A[i] *= A[i]
	}
	sort.Ints(A)
	return A
}

/*
执行用时 : 32 ms , 在所有 Go 提交中击败了 93.93% 的用户
内存消耗 : 6.4 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
