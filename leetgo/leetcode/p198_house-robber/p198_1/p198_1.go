package p198_1

// 纯递归效率较低
// Deprecated
func Rob(nums []int) int {
	if nums == nil || len(nums) < 1 {
		return 0
	}
	a = nums
	return f(len(a) - 1)
}

var a []int

func f(x int) int {
	switch x {
	case 0:
		return a[0]
	case 1:
		return max(a[0], a[1])
	case 2:
		return max(a[1], a[0]+a[2])
	default:
		return max(max(f(x-2), f(x-3))+a[x], f(x-1))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
