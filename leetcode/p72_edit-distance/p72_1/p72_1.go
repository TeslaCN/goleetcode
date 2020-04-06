package p72_1

func minDistance(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)
	switch {
	case word1 == word2:
		return 0
	case n1 == 0:
		return n2
	case n2 == 0:
		return n1
	}
	cache := make([][]int, n1+1)

	for i := 0; i <= n1; i++ {
		cache[i] = make([]int, n2+1)
	}
	for i := 0; i <= n1; i++ {
		cache[i][0] = i
	}
	for j := 0; j <= n2; j++ {
		cache[0][j] = j
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			cache[i][j] = min(cache[i-1][j]+1, cache[i][j-1]+1)
			if word1[i-1] != word2[j-1] {
				cache[i][j] = min(cache[i][j], cache[i-1][j-1]+1)
			} else {
				cache[i][j] = min(cache[i][j], cache[i-1][j-1])
			}
		}
	}
	return cache[n1][n2]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
执行用时 : 4 ms , 在所有 Go 提交中击败了 84.65% 的用户
内存消耗 : 5.7 MB , 在所有 Go 提交中击败了 80.00% 的用户
*/
