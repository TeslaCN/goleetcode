package p820_3

import (
	"sort"
	"strings"
)

var MinimumLengthEncoding = minimumLengthEncoding

func minimumLengthEncoding(words []string) int {
	if len(words) <= 1 {
		return len(append(words, "")[0]) + 1
	}
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	length := 0
	for i := 0; i < len(words)-1; i++ {
		longest := true
		for j := i + 1; j < len(words); j++ {
			if strings.HasSuffix(words[j], words[i]) {
				longest = false
				break
			}
		}
		if longest {
			length += len(words[i]) + 1
		}
	}
	return length + len(words[len(words)-1]) + 1
}

/*
执行用时 : 224 ms , 在所有 Go 提交中击败了 28.30% 的用户
内存消耗 : 6 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
