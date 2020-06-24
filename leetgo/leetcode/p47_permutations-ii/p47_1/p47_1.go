package p47_1

import "sort"

var PermuteUnique = permuteUnique

func permuteUnique(nums []int) [][]int {
	if len(nums) < 1 {
		return [][]int{}
	}
	sort.Ints(nums)
	var sets [][]int
	var p func([]int, int)

	p = func(nums []int, from int) {
		if from+1 == len(nums) {
			copied := make([]int, len(nums))
			copy(copied, nums)
			sets = append(sets, copied)
			return
		}
		for i := from; i < len(nums); i++ {
			if i != from && nums[i] == nums[from] {
				continue
			}
			nums[from], nums[i] = nums[i], nums[from]
			copied := make([]int, len(nums))
			copy(copied, nums)
			p(copied, from+1)
		}
	}
	p(nums, 0)
	return sets
}
