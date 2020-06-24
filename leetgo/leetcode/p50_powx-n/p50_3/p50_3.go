package p50_3

var MyPow = myPow

// Runtime: O(log(n))
// Space: O(1)
func myPow(x float64, n int) float64 {
	negative := n < 0
	if negative {
		n = -n
	}
	r := 1.0
	cx := x
	for n > 0 {
		if n&1 == 1 {
			r *= cx
		}
		cx *= cx
		n >>= 1
	}
	if negative {
		return 1.0 / r
	}
	return r
}

/*
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 2 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
