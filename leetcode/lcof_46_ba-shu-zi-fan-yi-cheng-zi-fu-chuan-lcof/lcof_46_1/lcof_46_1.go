package lcof_46_1

var TranslateNum = translateNum

// Runtime: O(log(n))
// Space: O(log(n))
func translateNum(num int) int {
	if num == 0 {
		return 1
	}
	if mod := num % 100; mod >= 10 && mod <= 25 {
		return translateNum(num/10) + translateNum(num/100)
	} else {
		return translateNum(num / 10)
	}
}

/*
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 1.9 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
