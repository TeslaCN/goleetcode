package p153_1

// binary search
// Runtime: O(log(n))
// Space: O(1)
func findMin(nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	}
	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}

	left, right := 0, len(nums)-1

	for left+1 < right {
		mid := (left + right) >> 1
		if nums[mid] < nums[left] {
			right = mid
		} else {
			left = mid
		}
	}
	return nums[right]
}
