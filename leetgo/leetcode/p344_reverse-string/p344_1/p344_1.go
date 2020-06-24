package p344_1

func reverseString(s []byte) {
	n := len(s)
	for i := 0; i < n>>1; i++ {
		s[i], s[n-1-i] = s[n-1-i], s[i]
	}
}

/*
执行用时 : 44 ms , 在所有 Go 提交中击败了 73.97% 的用户
内存消耗 : 6.3 MB , 在所有 Go 提交中击败了 33.33% 的用户
*/
