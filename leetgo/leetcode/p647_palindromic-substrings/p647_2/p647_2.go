package p647_2

var CountSubstrings = countSubstrings

func countSubstrings(s string) int {
	n, ans := len(s), 0
	if n == 0 {
		return ans
	}
	for step := 0; step <= 1; step++ {
		for cur := 0; cur < n-step; cur++ {
			for left, right := cur, cur+step; left >= 0 && right < n; {
				if s[left] == s[right] {
					ans++
				} else {
					break
				}
				left--
				right++
			}
		}
	}
	return ans
}

/*
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 1.9 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
