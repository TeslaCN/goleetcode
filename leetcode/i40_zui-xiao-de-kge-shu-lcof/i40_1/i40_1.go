package i40_1

import "sort"

var GetLeastNumbers = getLeastNumbers

func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}
