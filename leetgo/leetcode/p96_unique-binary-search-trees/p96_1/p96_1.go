package p96_1

var NumTrees = numTrees

// Space: O(1)
func numTrees(n int) int {
	if n < 2 {
		return 1
	}
	sum := 0
	for i := 0; i < n>>1; i++ {
		sum += 2 * numTrees(i) * numTrees(n-1-i)
	}
	if n&1 == 1 {
		sum += numTrees(n>>1) * numTrees(n>>1)
	}
	return sum
}

/*
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 1.9 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
