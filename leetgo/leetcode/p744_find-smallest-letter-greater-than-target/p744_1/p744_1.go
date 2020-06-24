package p744_1

// binary search
// Runtime: O(log(n))
// Space: O(1)
func nextGreatestLetter(letters []byte, target byte) byte {
	switch len(letters) {
	case 0:
		return 0
	case 1:
		return letters[0]
	}
	if target >= letters[len(letters)-1] || target < letters[0] {
		return letters[0]
	}
	left, right := 0, len(letters)-1
	for left < right {
		mid := (left + right) >> 1
		if letters[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return letters[left]
}
