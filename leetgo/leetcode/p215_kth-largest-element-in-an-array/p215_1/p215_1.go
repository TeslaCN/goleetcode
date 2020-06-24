package p215_1

var FindKthLargest = findKthLargest

// Runtime: O(nlog(n))
// Space: O(1)
func findKthLargest(nums []int, k int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	i := partition(nums)
	if len(nums)-i-1 >= k {
		return findKthLargest(nums[i+1:], k)
	} else {
		return findKthLargest(nums[:i+1], k-len(nums)+i+1)
	}
}

func partition(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	left, right := 0, len(nums)-1
	base := nums[0]
	for left < right {
		for left < right && nums[right] >= base {
			right--
		}
		for left < right && nums[left] <= base {
			left++
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	nums[0] = nums[left]
	nums[left] = base
	return left
}
