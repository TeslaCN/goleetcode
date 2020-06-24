package i08_11_2

// Deprecated
var WaysToChange = waysToChange

const mod = 1000000007

// Deprecated
func waysToChange(n int) int {
	n /= 5
	if n < 2 {
		return 1 + n
	}
	sum := 4
	for i := 2; i < n; i++ {
		sum = (sum + i) % mod
	}
	return sum
}
