package p169_1

// Runtime: O(n)
// Space: O(n)
func MajorityElement(nums []int) int {
	if nums == nil {
		return 0
	}
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		if m[nums[i]] > len(nums)/2 {
			return nums[i]
		}
	}
	return 0
}
