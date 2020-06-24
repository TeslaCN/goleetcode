package p11_1

var MaxArea = maxArea

// Runtime: O(n^2)
// Space: O(1)
func maxArea(height []int) int {
	if height == nil || len(height) < 2 {
		return 0
	}
	maxVal, length := 0, len(height)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			maxVal = max(maxVal, (j-i)*min(height[i], height[j]))
		}
	}
	return maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
