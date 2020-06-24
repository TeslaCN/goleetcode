package p338_3

var CountBits = countBits

// Runtime: O(n)
func countBits(num int) []int {
	if num < 0 {
		return []int{}
	}
	counts := make([]int, num+1)
	counts[0] = 0
	for i := 0; i < len(counts); i++ {
		counts[i] = i&1 + counts[i>>1]
	}
	return counts
}
