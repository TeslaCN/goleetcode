package p5_1

func longestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	maxLen, maxLeft, maxRight := 0, -1, -1
	for step := 0; step <= 1; step++ {
		for i := 0; i < n-step; i++ {
			for left, right := i, i+step; left >= 0 && right < n; {
				if s[left] != s[right] {
					break
				}
				if l := right - left + 1; l > maxLen {
					maxLen, maxLeft, maxRight = l, left, right
				}
				left--
				right++
			}
		}
	}
	return s[maxLeft : maxRight+1]
}

/*
执行用时 : 8 ms , 在所有 Go 提交中击败了 72.65% 的用户
内存消耗 : 2.2 MB , 在所有 Go 提交中击败了 44.83% 的用户
*/
