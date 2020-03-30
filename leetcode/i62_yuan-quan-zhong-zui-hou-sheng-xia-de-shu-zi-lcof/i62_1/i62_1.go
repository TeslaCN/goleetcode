package i62_1

// Deprecated
var LastRemaining = lastRemaining

// Deprecated
func lastRemaining(n int, m int) int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i
	}
	cur := n - 1
	remain := n
	for remain > 1 {
		for i := 0; i < m; i++ {
			cur = (cur + 1) % n
			for nums[cur] == -1 {
				cur = (cur + 1) % n
			}
		}
		nums[cur] = -1
		remain--
	}
	for i := 0; i < n; i++ {
		if nums[i] > -1 {
			return nums[i]
		}
	}
	return 0
}
