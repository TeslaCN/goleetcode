package p56_1

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
	})
	for i := 0; i < len(intervals); i++ {
		start, end := intervals[i][0], intervals[i][1]
		j := i + 1
		for ; j < len(intervals) && intervals[j][0] <= end; j++ {
			if intervals[j][1] > end {
				end = intervals[j][1]
			}
		}
		intervals[j-1][0], intervals[j-1][1] = start, end
		intervals = append(intervals[:i], intervals[j-1:]...)
	}
	return intervals
}

/*
执行用时 : 8 ms , 在所有 Go 提交中击败了 97.71% 的用户
内存消耗 : 4.6 MB , 在所有 Go 提交中击败了 100.00% 的用户

执行用时 : 12 ms , 在所有 Go 提交中击败了 79.14% 的用户
内存消耗 : 4.6 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
