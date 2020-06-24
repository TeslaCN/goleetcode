package p912_quicksort

var SortArray = sortArray

func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	pivot := nums[0]
	left, right := 0, len(nums)-1
	for left < right {
		for left < right && nums[right] >= pivot {
			right--
		}
		for left < right && nums[left] <= pivot {
			left++
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	nums[left], nums[0] = nums[0], nums[left]
	sortArray(nums[:left])
	sortArray(nums[left+1:])
	return nums
}

/*
执行用时 : 20 ms , 在所有 Go 提交中击败了 99.07% 的用户
内存消耗 : 6.5 MB , 在所有 Go 提交中击败了 5.13% 的用户
*/
