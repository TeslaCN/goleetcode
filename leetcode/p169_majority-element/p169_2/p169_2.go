package p169_2

// Space: O(1)
// Runtime: nlog(n)
// 摩尔投票
func MajorityElement(nums []int) int {
	if nums == nil || len(nums) < 1 {
		return 0
	}
	count, target := 0, nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == target {
			count++
		} else {
			count--
		}
		if count < 0 && i < len(nums)-1 {
			i++
			target = nums[i]
			count = 0
		}
	}
	return target
}
