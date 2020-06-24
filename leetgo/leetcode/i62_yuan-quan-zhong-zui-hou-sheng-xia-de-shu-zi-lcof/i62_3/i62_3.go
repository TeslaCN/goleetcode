package i62_3

var LastRemaining = lastRemaining

func lastRemaining(n int, m int) int {
	switch {
	case n == 1:
		return 0
	case m == 1:
		return n - 1
	}
	num := 0
	for i := 2; i <= n; i++ {
		num += m
		if num >= i {
			num %= i
		}
	}
	return num
}
