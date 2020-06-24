package p9_1

import "math"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	length, reversed := int(math.Log10(float64(x)))+1, 0
	for i := 0; i < length/2; i++ {
		reversed, x = reversed*10+(x%10), x/10
	}
	if length&1 == 1 {
		x /= 10
	}
	return x == reversed
}

/*
执行用时 : 8 ms , 在所有 Go 提交中击败了 97.81% 的用户
内存消耗 : 5.2 MB , 在所有 Go 提交中击败了 88.00% 的用户
*/
