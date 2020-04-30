package p560_1

var SubarraySum = subarraySum

// Runtime: O(n^2)
// Space: O(1)
func subarraySum(nums []int, k int) int {
	hits := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				hits++
			}
		}
		sum = 0
	}
	return hits
}

/*
执行用时 : 212 ms , 在所有 Go 提交中击败了 25.50% 的用户
内存消耗 : 5.2 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
