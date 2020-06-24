package p912_shellsort

var SortArray = sortArray

func sortArray(nums []int) []int {
	for gap := len(nums) >> 1; gap > 0; gap >>= 1 {
		for i := 0; i < gap; i++ {
			for j := gap + i; j < len(nums); j += gap {
				e := nums[j]
				k := j - gap
				for ; k >= 0; k -= gap {
					if nums[k] > e {
						nums[k+gap] = nums[k]
					} else {
						break
					}
				}
				nums[k+gap] = e
			}
		}
	}
	return nums
}

/*
执行用时 : 28 ms , 在所有 Go 提交中击败了 58.51% 的用户
内存消耗 : 6.5 MB , 在所有 Go 提交中击败了 5.13% 的用户
*/
