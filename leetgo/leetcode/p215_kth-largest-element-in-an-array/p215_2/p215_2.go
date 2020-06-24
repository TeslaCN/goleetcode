package p215_2

var FindKthLargest = findKthLargest

// same as p215_1
func findKthLargest(nums []int, k int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	base := nums[0]
	left, right := 0, len(nums)-1
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
	if len(nums)-left-1 >= k {
		return findKthLargest(nums[left+1:], k)
	} else {
		return findKthLargest(nums[:left+1], k-(len(nums)-left-1))
	}
}
