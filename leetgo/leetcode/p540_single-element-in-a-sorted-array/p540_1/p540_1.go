package p540_1

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
		switch mid & 1 {
		case 0:
			switch nums[mid] {
			case nums[mid+1]:
				left = mid
			case nums[mid-1]:
				right = mid
			default:
				return nums[mid]
			}
		case 1:
			switch nums[mid] {
			case nums[mid+1]:
				right = mid - 1
			case nums[mid-1]:
				left = mid + 1
			default:
				return nums[mid]
			}
		}
	}
	return nums[left]
}
