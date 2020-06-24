package p763_1

import "strings"

func partitionLabels(S string) []int {
	if len(S) <= 1 {
		return []int{len(S)}
	}
	partitions := make([]int, 0)
	from := 0
	for from < len(S) {
		maxIndex := strings.LastIndexByte(S, S[from])
		for i := from + 1; i < maxIndex; i++ {
			index := strings.LastIndexByte(S, S[i])
			if index > maxIndex {
				maxIndex = index
			}
		}
		partitions = append(partitions, maxIndex+1-from)
		from = maxIndex + 1
	}
	return partitions
}
