package p647_1

var CountSubstrings = countSubstrings

// Brute Force
func countSubstrings(s string) int {
	n := len(s)
	if s == "" {
		return 0
	}
	ans := n
	for step := 2; step <= n; step++ {
		for i := 0; i <= n-step; i++ {
			if isPalindrome(s[i : i+step]) {
				ans++
			}
		}
	}
	return ans
}

func isPalindrome(s string) bool {
	if s == "" {
		return false
	}
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

/*
执行用时 : 64 ms , 在所有 Go 提交中击败了 17.02% 的用户
内存消耗 : 2 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
