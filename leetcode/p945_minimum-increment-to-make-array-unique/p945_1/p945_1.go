package p945_1

var MinIncrementForUnique = minIncrementForUnique

// Deprecated: Too slow
func minIncrementForUnique(A []int) int {
	if len(A) <= 1 {
		return 0
	}
	result := 0
	m := map[int]int{}
	var toAdd []int
	for _, value := range A {
		if m[value] == 0 {
			m[value]++
		} else {
			toAdd = append(toAdd, value)
		}
	}
	for _, value := range toAdd {
		newValue := value + 1
		for ; m[newValue] > 0; newValue++ {
		}
		m[newValue] = 1
		result += newValue - value
	}
	return result
}
