package p11_2

var MaxArea = maxArea

// Runtime: O(n)
// Space: O(1)
func maxArea(height []int) int {
	if height == nil || len(height) < 2 {
		return 0
	}
	left, right := 0, len(height)-1
	maxVal := 0
	for left < right {
		maxVal = max(maxVal, areaOf(left, right, height))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxVal
}

func areaOf(l, r int, height []int) int {
	return min(height[l], height[r]) * (r - l)
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
