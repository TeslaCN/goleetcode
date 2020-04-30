package p139_2

import "strings"

var WordBreak = wordBreak

func wordBreak(s string, wordDict []string) bool {
	if s == "" {
		return true
	} else if len(wordDict) < 1 {
		return false
	}
	cache := make([]bool, len(s)+1)
	cache[0] = true
	minLen := len(wordDict[0])
	for i := 1; i < len(wordDict); i++ {
		if len(wordDict[i]) < minLen {
			minLen = len(wordDict[i])
		}
	}
	for i := 1; i < minLen; i++ {
		cache[i] = false
	}
	for i := minLen; i <= len(s); i++ {
		for _, word := range wordDict {
			if len(word) <= i && strings.HasSuffix(s[:i], word) && cache[i-len(word)] {
				cache[i] = true
				break
			}
		}
	}
	return cache[len(s)]
}

/*
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 2.1 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
