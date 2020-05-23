package p85_1

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) <= 0 || len(matrix[0]) <= 0 {
		return 0
	}
	matrixHeights := make([][]int, len(matrix))
	matrixHeights[0] = make([]int, len(matrix[0]))
	for i := 0; i < len(matrix[0]); i++ {
		matrixHeights[0][i] = int(matrix[0][i] - '0')
	}
	for i := 1; i < len(matrixHeights); i++ {
		matrixHeights[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrixHeights[i]); j++ {
			if matrix[i][j] == '1' {
				matrixHeights[i][j] = matrixHeights[i-1][j] + 1
			}
		}
	}
	maxArea := 0
	for _, heights := range matrixHeights {
		if area := maxAreaInHistogram(heights); area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

func maxAreaInHistogram(heights []int) int {
	stack, maxArea := []int{-1}, 0
	for i, height := range heights {
		if len(stack) == 1 || height >= heights[stack[len(stack)-1]] {
			stack = append(stack, i)
			continue
		}
		for len(stack) > 1 && height < heights[stack[len(stack)-1]] {
			if area := (i - 1 - stack[len(stack)-2]) * heights[stack[len(stack)-1]]; area > maxArea {
				maxArea = area
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for len(stack) > 1 {
		if area := (len(heights) - 1 - stack[len(stack)-2]) * heights[stack[len(stack)-1]]; area > maxArea {
			maxArea = area
		}
		stack = stack[:len(stack)-1]
	}
	return maxArea
}

/*
执行用时 : 4 ms , 在所有 Go 提交中击败了 84.88% 的用户
内存消耗 : 4.7 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
