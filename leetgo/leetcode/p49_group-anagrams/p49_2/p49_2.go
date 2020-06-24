package p49_2

var GroupAnagrams = groupAnagrams

var hash = [26]int{
	2, 3, 5, 7, 11,
	13, 17, 19, 23, 29,
	31, 37, 41, 43, 47,
	53, 59, 61, 67, 71,
	73, 79, 83, 89, 97, 101,
}

func groupAnagrams(strs []string) [][]string {
	cache := map[int][]string{}
	for _, str := range strs {
		sum := 1
		for i := 0; i < len(str); i++ {
			sum *= hash[str[i]-0x61]
		}
		if _, exists := cache[sum]; !exists {
			cache[sum] = []string{str}
		} else {
			cache[sum] = append(cache[sum], str)
		}
	}
	var result [][]string
	for _, strings := range cache {
		result = append(result, strings)
	}
	return result
}

/*
执行用时 : 20 ms , 在所有 Go 提交中击败了 98.39% 的用户
内存消耗 : 6.7 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
