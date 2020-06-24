package p76_1

func minWindow(s string, t string) string {
	if s == "" || t == "" || len(s) < len(t) {
		return ""
	}
	from, end, cache, totalRemain, minFrom, minEnd := 0, 0, map[byte]int{}, 0, -1, -1
	for i := 0; i < len(t); i++ {
		cache[t[i]]++
		totalRemain++
	}
	for end < len(s) || totalRemain == 0 {
		for from < end && totalRemain == 0 {
			if minFrom == -1 || end-from < minEnd-minFrom {
				minFrom, minEnd = from, end
			}
			if remain, exists := cache[s[from]]; exists {
				cache[s[from]]++
				if remain >= 0 {
					totalRemain++
				}
			}
			from++
		}
		for end < len(s) && totalRemain > 0 {
			end++
			if remain, exists := cache[s[end-1]]; exists {
				cache[s[end-1]]--
				if remain > 0 {
					totalRemain--
				}
			}
		}
	}
	if minFrom < 0 || minEnd < 0 {
		return ""
	}
	return s[minFrom:minEnd]
}

/*
执行用时 : 20 ms , 在所有 Go 提交中击败了 78.82% 的用户
内存消耗 : 2.9 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
