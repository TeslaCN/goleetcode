package p128_2

var LongestConsecutive = longestConsecutive

// Runtime: O(n)
// Space: O(n)
func longestConsecutive(nums []int) int {
	cache := map[int]bool{}
	for _, num := range nums {
		cache[num] = true
	}
	if len(cache) == 0 {
		return 0
	}
	maxLen := 1
	for _, num := range nums {
		if !cache[num-1] {
			l := 0
			for i := num; cache[i]; i++ {
				l++
			}
			if l > maxLen {
				maxLen = l
			}
		}
	}
	return maxLen
}

/*
执行用时 : 8 ms , 在所有 Go 提交中击败了 68.84% 的用户
内存消耗 : 3.6 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
