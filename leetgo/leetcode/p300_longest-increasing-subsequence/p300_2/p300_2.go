package p300_2

var LengthOfLIS = lengthOfLIS

// Runtime: O(nlog(n))
// Space: O(n)
func lengthOfLIS(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	maxL := 0
	sub := make([]int, len(nums))
	for _, num := range nums {
		left, right := 0, maxL
		for left < right {
			mid := left + (right-left)>>1
			if sub[mid] < num {
				left = mid + 1
			} else {
				right = mid
			}
		}
		sub[left] = num
		if left == maxL {
			maxL++
		}
	}
	return maxL
}
