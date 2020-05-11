package p50_2

var MyPow = myPow

// Runtime: O(log(n))
// Space: O(log(n))
func myPow(x float64, n int) float64 {
	switch {
	case n == 0:
		return 1
	case n == 1:
		return x
	case n == -1:
		return 1 / x
	case n&1 == 0:
		result := myPow(x, n>>1)
		return result * result
	default:
		return myPow(x, n>>1) * myPow(x, n-n>>1)
	}
}

/*
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 2.1 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
