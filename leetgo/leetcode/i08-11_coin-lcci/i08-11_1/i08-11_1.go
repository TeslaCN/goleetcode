package i08_11_1

// Deprecated
var WaysToChange = waysToChange

// Deprecated
func waysToChange(n int) int {
	return change(n/5, []int{5, 2, 1})
}

func change(target int, options []int) int {
	if target == 0 || len(options) < 1 {
		return 1
	}
	sum := 0
	for i := 0; i <= target/options[0]; i++ {
		sum += change(target-options[0]*i, options[1:])
	}
	return sum
}
