package p820_1

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
	from, to := 0, 0
	for from < len(words) {
		toDelete := false
		for i := from + 1; i < len(words); i++ {
			if strings.HasSuffix(words[i], words[to]) {
				toDelete = true
				from++
				break
			}
		}
		words[to] = words[from]
		if !toDelete {
			to++
			from++
		}
	}
	length := 0
	for i := 0; i < to; i++ {
		length += len(words[i]) + 1
	}
	return length
}

/*
执行用时 : 232 ms , 在所有 Go 提交中击败了 26.41% 的用户
内存消耗 : 5.9 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
