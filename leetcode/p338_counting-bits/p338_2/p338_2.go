package p338_2

import "math"

var CountBits = countBits

// Runtime: O(n)?
func countBits(num int) []int {
	if num < 0 {
		return []int{}
	}
	counts := make([]int, num+1)
	counts[0] = 0
	for i := 1; i < len(counts); i++ {
		number := i
		index := math.Log2(float64(number))
		intIndex := int(index)
		if index == float64(intIndex) {
			counts[i] = 1
		} else {
			number -= int(math.Pow(2, float64(intIndex)))
			counts[i] = counts[number] + 1
		}
	}
	return counts
}
