package p974_1

var SubarraysDivByK = subarraysDivByK

// TLE
// Space: O(n)
// Runtime: O(n^2)
// Deprecated
func subarraysDivByK(A []int, K int) int {
	n := len(A)
	if n == 0 {
		return 0
	}
	cache, ans := make([]int, n), 0
	cache[0] = A[0]
	for i := 1; i < n; i++ {
		cache[i] = cache[i-1] + A[i]
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if (cache[i]-cache[j])%K == 0 {
				ans++
			}
		}
		if cache[i]%K == 0 {
			ans++
		}
	}
	return ans
}
