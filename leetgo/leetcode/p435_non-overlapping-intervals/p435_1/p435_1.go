package p435_1

import (
	"sort"
)

var EraseOverlapIntervals = eraseOverlapIntervals

type IntervalInterface [][]int

func (o IntervalInterface) Len() int {
	return len(o)
}
func (o IntervalInterface) Less(i, j int) bool {
	return o[i][1] < o[j][1]

}
func (o IntervalInterface) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 0
	}
	o := IntervalInterface(intervals)
	sort.Sort(o)
	keep := 1
	pre := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= pre {
			keep++
			pre = intervals[i][1]
		}
	}
	return len(intervals) - keep
}
