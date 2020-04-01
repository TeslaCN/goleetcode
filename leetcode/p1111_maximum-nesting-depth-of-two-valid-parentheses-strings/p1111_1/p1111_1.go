package p1111_1

func maxDepthAfterSplit(seq string) []int {
	answer := make([]int, len(seq))
	a, b := 0, 0
	for i := 0; i < len(seq); i++ {
		switch seq[i] {
		case '(':
			if a <= b {
				a++
				answer[i] = 0
			} else {
				b++
				answer[i] = 1
			}
		case ')':
			if a > b {
				answer[i] = 0
				a--
			} else {
				answer[i] = 1
				b--
			}
		default:
		}
	}
	return answer
}

/*
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 2.9 MB , 在所有 Go 提交中击败了 25.00% 的用户
*/
