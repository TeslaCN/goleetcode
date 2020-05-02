package p3_1

import "math"

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n < 2 {
		return n
	}
	max, from, end := 1, 0, 0
	m := [math.MaxUint8]int{}
	for end < n {
		for m[s[end]] == 1 {
			m[s[from]]--
			from++
		}
		for end < n && m[s[end]] == 0 {
			m[s[end]]++
			end++
		}
		if l := end - from; l > max {
			max = l
		}
	}
	return max
}

/*
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 2.6 MB , 在所有 Go 提交中击败了 87.10% 的用户
*/
