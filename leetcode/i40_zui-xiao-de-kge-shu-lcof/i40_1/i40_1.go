package i40_1

import "sort"

var GetLeastNumbers = getLeastNumbers

func getLeastNumbers(arr []int, k int) []int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	return arr[len(arr)-k:]
}
