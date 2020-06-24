package p912_golang

import "sort"

var SortArray = sortArray

func sortArray(nums []int) []int {
	sort.Ints(nums)
	return nums
}

/*
执行用时 : 28 ms , 在所有 Go 提交中击败了 58.51% 的用户
内存消耗 : 6.4 MB , 在所有 Go 提交中击败了 5.13% 的用户
*/
