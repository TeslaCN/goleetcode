package p665_1

var CheckPossibility = checkPossibility

// Runtime: O(n)
// Space: O(1)
func checkPossibility(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}
	return isNonDecreasing(nums, true)
}

func isNonDecreasing(nums []int, modifiable bool) bool {
	if len(nums) <= 1 {
		return true
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if modifiable {
				tmp := nums[i]
				nums[i] = nums[i+1]
				res := isNonDecreasing(nums, false)
				nums[i] = tmp
				tmp = nums[i+1]
				nums[i+1] = nums[i]
				res = res || isNonDecreasing(nums, false)
				return res
			} else {
				return false
			}
		}
	}
	return true
}
