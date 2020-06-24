package p139_1

import "strings"

// Deprecated: TLE
var WordBreak = wordBreak

// Deprecated: TLE
func wordBreak(s string, wordDict []string) bool {
	if s == "" {
		return true
	} else if len(wordDict) < 1 {
		return false
	}
	for _, word := range wordDict {
		if strings.HasPrefix(s, word) && wordBreak(strings.TrimPrefix(s, word), wordDict) {
			return true
		}
	}
	return false
}
