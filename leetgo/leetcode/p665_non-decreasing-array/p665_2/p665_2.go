package p665_2

var CheckPossibility = checkPossibility

// Runtime: O(n)
// Space: O(1)
func checkPossibility(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}
	remain := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			if remain <= 0 {
				return false
			}
			if i-2 < 0 || nums[i] >= nums[i-2] {
				nums[i-1] = nums[i]
			} else {
				nums[i] = nums[i-1]
			}
			remain--
		}
	}
	return true
}
