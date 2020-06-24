package p78_1

import "sort"

var Subsets = subsets

func subsets(nums []int) [][]int {
	if nums == nil {
		return [][]int{}
	}
	sort.Ints(nums)
	sets := [][]int{{}}
	for i := 0; i < len(nums); i++ {
		sets = append(sets, sub([]int{}, nums, i+1)...)
	}
	return sets
}

func sub(prefix, options []int, need int) [][]int {
	need--
	var sets [][]int
	for i := 0; i < len(options); i++ {
		// todo take notes
		copied := make([]int, len(prefix))
		copy(copied, prefix)
		if need > 0 {
			sets = append(sets, sub(append(copied, options[i]), options[i+1:], need)...)
		} else {
			sets = append(sets, append(copied, options[i]))
		}
	}
	return sets
}
