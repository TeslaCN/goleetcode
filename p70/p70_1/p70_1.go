package p70_1

// Deprecated: Timeout
func ClimbStairs(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return ClimbStairs(n-1) + ClimbStairs(n-2)
}
