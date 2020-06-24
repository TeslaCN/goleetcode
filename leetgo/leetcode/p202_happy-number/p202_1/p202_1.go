package p202_1

func isHappy(n int) bool {
	cache := map[int]struct{}{}
	sum := 0
	for {
		for n > 0 {
			base := n % 10
			sum += base * base
			n /= 10
		}
		if sum == 1 {
			break
		}
		if _, exists := cache[sum]; exists {
			return false
		}
		cache[sum] = struct{}{}
		n, sum = sum, 0
	}
	return true
}

/*
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 2.1 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
