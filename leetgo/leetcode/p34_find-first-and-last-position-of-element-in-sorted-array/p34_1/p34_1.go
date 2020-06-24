package p34_1

var SearchRange = searchRange

// Runtime: O(log(n))
func searchRange(nums []int, target int) []int {
	if nums == nil || len(nums) < 1 {
		return []int{-1, -1}
	}
	return []int{bSearch(nums, 0, len(nums)-1, target, false), bSearch(nums, 0, len(nums)-1, target, true)}
}

func bSearch(nums []int, left, right, target int, max bool) int {
	latestIndex := -1
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			latestIndex = mid
			if max {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	if nums[left] == target {
		return left
	}
	return latestIndex
}
