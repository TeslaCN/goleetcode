package p15_1

import (
	"sort"
)

var ThreeSum = threeSum

// Runtime: O(n^2 * log(n))
func threeSum(nums []int) [][]int {
	if nums == nil || len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	tuples := make([][]int, 0)
	result := make(map[[3]int]bool)
	for i := 0; i < len(nums)-2 && nums[i] <= 0; i++ {

		expect2Sum := -nums[i]
		for j := i + 1; j < len(nums)-1; j++ {
			expectNum := expect2Sum - nums[j]

			if nums[j+1] > expectNum {
				break
			}

			index := sort.SearchInts(nums, expectNum)
			if index < len(nums) && nums[index] == expectNum {
				result[[3]int{nums[i], nums[j], nums[index]}] = true
			}
		}
	}
	for key := range result {
		tuples = append(tuples, []int{key[0], key[1], key[2]})
	}
	return tuples
}
