package p96_2

var NumTrees = numTrees

// Space: O(n)
func numTrees(n int) int {
	if n < 2 {
		return 1
	}
	cache := make([]int, n+1)
	cache[0], cache[1], cache[2] = 1, 1, 2
	for i := 3; i <= n; i++ {
		for j := 0; j < i>>1; j++ {
			cache[i] += 2 * cache[j] * cache[i-1-j]
		}
		if i&1 == 1 {
			cache[i] += cache[i>>1] * cache[i>>1]
		}
	}
	return cache[n]
}

/*
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 2 MB , 在所有 Go 提交中击败了 66.67% 的用户
*/
