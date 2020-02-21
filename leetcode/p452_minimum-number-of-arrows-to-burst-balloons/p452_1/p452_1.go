package p452_1

import "sort"

var FindMinArrowShots = findMinArrowShots

type PointInterface [][]int

func (p PointInterface) Len() int {
	return len(p)
}

func (p PointInterface) Less(i, j int) bool {
	return p[i][1] < p[j][1]
}

func (p PointInterface) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Space: O(1)
// Runtime: O(n*log(n))
func findMinArrowShots(points [][]int) int {
	if len(points) < 2 {
		return len(points)
	}
	sort.Sort(PointInterface(points))
	count := 1
	pre := points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] > pre {
			pre = points[i][1]
			count++
		}
	}
	return count
}
