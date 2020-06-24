package p912_insertsort

var SortArray = sortArray

func sortArray(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		e := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			if nums[j] > e {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[j+1] = e
	}
	return nums
}

/*
执行用时 : 428 ms , 在所有 Go 提交中击败了 10.72% 的用户
内存消耗 : 6.4 MB , 在所有 Go 提交中击败了 5.13% 的用户
*/
