package p455_2

import "sort"

var FindContentChildren = findContentChildren

// Runtime: O(n*log(n))
// Space: O(1)
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	count := 0
	pk, pb := 0, 0
	for pk < len(g) && pb < len(s) {
		if s[pb] >= g[pk] {
			pk++
			count++
		}
		pb++
	}
	return count
}
