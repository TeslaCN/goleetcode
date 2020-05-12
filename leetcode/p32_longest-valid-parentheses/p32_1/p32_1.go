package p32_1

// Runtime: O(n)
// Space: O(1)
func longestValidParentheses(s string) int {
	n := len(s)
	bracketsLeft, bracketsRight, lenLeft, lenRight, maxLeft, maxRight := 0, 0, 0, 0, 0, 0
	for i := 0; i < n; i++ {
		switch s[i] {
		case '(':
			bracketsLeft++
		case ')':
			bracketsLeft--
			if bracketsLeft >= 0 {
				lenLeft += 2
			} else {
				bracketsLeft = 0
				lenLeft = 0
			}
			if bracketsLeft == 0 {
				maxLeft = max(maxLeft, lenLeft)
			}
		}
		switch s[n-1-i] {
		case ')':
			bracketsRight++
		case '(':
			bracketsRight--
			if bracketsRight >= 0 {
				lenRight += 2
			} else {
				bracketsRight = 0
				lenRight = 0
			}
			if bracketsRight == 0 {
				maxRight = max(maxRight, lenRight)
			}
		}
	}
	return max(maxLeft, maxRight) & -2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 2.3 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
