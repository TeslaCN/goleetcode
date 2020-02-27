package p287_2

var FindDuplicate = findDuplicate

// Runtime: O(n)
// Space: O(1)
func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for {
		fast = nums[nums[fast]]
		slow = nums[slow]
		if slow == fast {
			p := 0
			for nums[p] != nums[slow] {
				p = nums[p]
				slow = nums[slow]
			}
			return nums[p]
		}
	}
}
