package p455_1

import "math"

var FindContentChildren = findContentChildren

// Runtime: O(n^2)
// Space: O(1)
func findContentChildren(g []int, s []int) int {
	count := 0
	for b, biscuit := range s {
		if biscuit < 0 {
			continue
		}
		giveTo := -1
		minOverflow := math.MaxInt32
		for k, kid := range g {
			if kid < 0 || biscuit < kid {
				continue
			}
			overflow := biscuit - kid
			if overflow < minOverflow {
				giveTo = k
				minOverflow = overflow
			}
		}
		if giveTo > -1 {
			s[b] = -1
			g[giveTo] = -1
			count++
		}
	}
	return count
}
