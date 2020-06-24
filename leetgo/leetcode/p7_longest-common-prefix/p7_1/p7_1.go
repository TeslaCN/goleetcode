package p7_1

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	b, pos := &strings.Builder{}, 0
	for {
		var c byte
		if pos < len(strs[0]) {
			c = strs[0][pos]
		} else {
			return b.String()
		}
		for i := 1; i < len(strs); i++ {
			if pos >= len(strs[i]) || strs[i][pos] != c {
				return b.String()
			}
		}
		b.WriteByte(c)
		pos++
	}
}

/*
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 2.4 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
