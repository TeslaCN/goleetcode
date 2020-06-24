package p912_bucketsort

var SortArray = sortArray

// Runtime: O(n)
// Space: O(n)
func sortArray(nums []int) []int {
	bucket := [100001]int{}
	for _, num := range nums {
		bucket[num+50000]++
	}
	cur := 0
	for i := 0; i < len(bucket); i++ {
		for ; bucket[i] > 0; bucket[i]-- {
			nums[cur] = i - 50000
			cur++
		}
	}
	return nums
}

/*
执行用时 : 20 ms , 在所有 Go 提交中击败了 99.07% 的用户
内存消耗 : 7.6 MB , 在所有 Go 提交中击败了 5.13% 的用户
*/
