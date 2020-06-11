package lcof_46_2

var TranslateNum = translateNum

// Runtime: O(log(n))
// Space: O(1)
func translateNum(num int) int {
	sum, preNum, pre10, pre100 := 1, num%10, 1, 1
	num /= 10
	for ; num > 0; num /= 10 {
		mod := num % 10
		if mod == 1 || mod == 2 && preNum <= 5 {
			sum = pre10 + pre100
		} else {
			sum = pre10
		}
		pre10, pre100, preNum = sum, pre10, mod
	}
	return sum
}

/*
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 2 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
