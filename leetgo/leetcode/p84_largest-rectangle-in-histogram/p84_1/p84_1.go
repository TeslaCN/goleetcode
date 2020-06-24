package p84_1

func largestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return heights[0]
	}
	stack, maxArea := []int{-1}, 0
	for i := 0; i < n; i++ {
		if len(stack) == 1 || heights[i] >= heights[stack[len(stack)-1]] {
			stack = append(stack, i)
			continue
		}
		for len(stack) > 1 && heights[stack[len(stack)-1]] > heights[i] {
			if area := (i - 1 - stack[len(stack)-1-1]) * heights[stack[len(stack)-1]]; area > maxArea {
				maxArea = area
			}
			if len(stack) > 1 {
				stack = stack[:len(stack)-1]
			}
		}
		stack = append(stack, i)
	}
	for len(stack) > 1 {
		if area := (n - 1 - stack[len(stack)-1-1]) * heights[stack[len(stack)-1]]; area > maxArea {
			maxArea = area
		}
		if len(stack) > 1 {
			stack = stack[:len(stack)-1]
		}
	}
	return maxArea
}

/*
执行用时 : 8 ms , 在所有 Go 提交中击败了 96.90% 的用户
内存消耗 : 5.9 MB , 在所有 Go 提交中击败了 33.33% 的用户
*/
