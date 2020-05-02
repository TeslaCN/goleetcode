package p438_2

var FindAnagrams = findAnagrams

func findAnagrams(s string, p string) []int {
	sLen, pLen := len(s), len(p)
	if sLen < 1 || pLen < 1 || sLen < pLen {
		return []int{}
	}
	result := make([]int, 0, sLen-pLen)
	sMap, pMap := [26]int{}, [26]int{}
	for i := 0; i < pLen; i++ {
		pMap[p[i]-0x61]++
	}
	for i := 0; i < pLen; i++ {
		sMap[s[i]-0x61]++
	}
	if sMap == pMap {
		result = append(result, 0)
	}
	for i := 0; i < sLen-pLen; i++ {
		sMap[s[i]-0x61]--
		sMap[s[i+pLen]-0x61]++
		if sMap == pMap {
			result = append(result, i+1)
		}
	}
	return result
}

/*
执行用时 : 4 ms , 在所有 Go 提交中击败了 97.25% 的用户
内存消耗 : 4.6 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
