package p974_3

var SubarraysDivByK = subarraysDivByK

// Runtime: O(n)
// Space: O(K)
// slice faster than map
func subarraysDivByK(A []int, K int) int {
	if len(A) == 0 {
		return 0
	}
	cache, ans, sum := make([]int, K), 0, 0
	cache[0] = 1
	for _, num := range A {
		sum += num
		mod := sum % K
		if mod < 0 {
			mod += K
		}
		cache[mod]++
		ans += cache[mod] - 1
	}
	return ans
}

/*
执行用时 : 44 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 6.5 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
