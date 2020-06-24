package p49_1

var GroupAnagrams = groupAnagrams

func groupAnagrams(strs []string) [][]string {
	cache := map[[26]byte][]string{}
	for _, str := range strs {
		b := [26]byte{}
		for i := 0; i < len(str); i++ {
			b[str[i]-0x61]++
		}
		if _, exists := cache[b]; !exists {
			cache[b] = make([]string, 0, 1)
		}
		cache[b] = append(cache[b], str)
	}
	result := make([][]string, 0, len(cache))
	for _, strings := range cache {
		result = append(result, strings)
	}
	return result
}

/*
执行用时 : 24 ms , 在所有 Go 提交中击败了 92.96% 的用户
内存消耗 : 7.4 MB , 在所有 Go 提交中击败了 33.33% 的用户
*/
