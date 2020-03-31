package p912_heapsort

var SortArray = sortArray

func sortArray(nums []int) []int {
	for i := (len(nums) - 1) >> 1; i >= 0; i-- {
		adjustHeap(nums, i, len(nums))
	}
	for i := len(nums) - 1; i >= 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		adjustHeap(nums, 0, i)
	}
	return nums
}

func adjustHeap(nums []int, pos, end int) {
	value := nums[pos]
	for child := pos<<1 + 1; child < end; child = child<<1 + 1 {
		if child+1 < end && nums[child] < nums[child+1] {
			child++
		}
		if value < nums[child] {
			nums[pos] = nums[child]
			pos = child
		}
	}
	nums[pos] = value
}

/*
执行用时 : 24 ms , 在所有 Go 提交中击败了 90.44% 的用户
内存消耗 : 6.5 MB , 在所有 Go 提交中击败了 5.13% 的用户
*/
