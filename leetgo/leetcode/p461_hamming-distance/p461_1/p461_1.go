package p461_1

// Runtime: O(1)
// Space: O(1)
func HammingDistance(x int, y int) int {
	n := x ^ y
	count := 0
	for ; n != 0; n >>= 1 {
		count += n & 1
	}
	return count
}
