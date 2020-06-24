package p540_2

var SingleNonDuplicate = singleNonDuplicate

// Runtime: O(log(n))
// Space: O(1)
func singleNonDuplicate(nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	}
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) >> 1
		if mid&1 == 0 {
			if nums[mid] == nums[mid+1] {
				left = mid
			} else if nums[mid] == nums[mid-1] {
				right = mid
			} else {
				return nums[mid]
			}
		} else {
			if nums[mid] == nums[mid+1] {
				right = mid - 1
			} else if nums[mid] == nums[mid-1] {
				left = mid + 1
			} else {
				return nums[mid]
			}
		}
	}
	return nums[left]
}
