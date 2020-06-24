package lcof_64_2

var SumNums = sumNums

func sumNums(n int) int {
	ans := 0
	var sum func(int) bool
	sum = func(a int) bool {
		ans += a
		return a > 0 && sum(a-1)
	}
	sum(n)
	return ans
}

/*
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 2.6 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
