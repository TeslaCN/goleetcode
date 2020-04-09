package p62_2

var UniquePaths = uniquePaths

func uniquePaths(m int, n int) int {
	// C(m-1++n-1,m-1)
	if m > n {
		m, n = n, m
	}
	a, b := 1, 1
	for i := n; i <= m+n-2; i++ {
		a *= i
	}
	for i := 2; i <= m-1; i++ {
		b *= i
	}
	return a / b
}

/*
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 1.9 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
