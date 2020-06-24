package p1371_1

var cmap = [0x61 + 26]int{
	'a': 0b00001,
	'e': 0b00010,
	'i': 0b00100,
	'o': 0b01000,
	'u': 0b10000,
}

func findTheLongestSubstring(s string) int {
	n, ans, status, cache := len(s), 0, 0, [1 << 5]int{0}
	for i := 1; i < len(cache); i++ {
		cache[i] = -1
	}
	for i := 0; i < n; i++ {
		status ^= cmap[s[i]]
		if cache[status] >= 0 {
			if value := i + 1 - cache[status]; value > ans {
				ans = value
			}
		} else {
			cache[status] = i + 1
		}
	}
	return ans
}

/*
执行用时 : 16 ms , 在所有 Go 提交中击败了 96.97% 的用户
内存消耗 : 6.2 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
