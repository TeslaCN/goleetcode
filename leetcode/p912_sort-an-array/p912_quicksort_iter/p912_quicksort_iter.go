package p912_quicksort_iter

var SortArray = sortArray

func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	queue := make([]int, 0)
	queue = append(queue, 0, len(nums))
	for len(queue) > 0 {
		start, end := queue[0], queue[1]
		queue = queue[2:]
		left, right := start, end-1
		if left >= right {
			continue
		}
		pivot := nums[start]
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
		nums[start], nums[left] = nums[left], nums[start]
		queue = append(queue, start, left)
		queue = append(queue, left+1, end)
	}
	return nums
}

/*
执行用时 : 20 ms , 在所有 Go 提交中击败了 99.07% 的用户
内存消耗 : 6.7 MB , 在所有 Go 提交中击败了 5.13% 的用户
*/
