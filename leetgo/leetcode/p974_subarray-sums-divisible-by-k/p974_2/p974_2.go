package p974_2

var SubarraysDivByK = subarraysDivByK

// Runtime: O(n)
// Space: O(K)
func subarraysDivByK(A []int, K int) int {
	if len(A) == 0 {
		return 0
	}
	cache, ans, sum := map[int]int{0: 1}, 0, 0
	for i := 0; i < len(A); i++ {
		sum += A[i]
		mod := sum % K
		if mod < 0 {
			mod += K
		}
		cache[mod]++
		ans += cache[mod] - 1
	}
	return ans
}
