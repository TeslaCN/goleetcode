package p42_3

var Trap = trap

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	sum := 0
	maxIndex := 0
	for i, value := range height {
		if value >= height[maxIndex] {
			maxIndex = i
		}
	}
	for left := 1; left < maxIndex; left++ {
		if height[left] < height[left-1] {
			sum += height[left-1] - height[left]
			height[left] = height[left-1]
		}
	}
	for right := len(height) - 2; right > maxIndex; right-- {
		if height[right] < height[right+1] {
			sum += height[right+1] - height[right]
			height[right] = height[right+1]
		}
	}
	return sum
}

/*
执行用时 : 4 ms , 在所有 Go 提交中击败了 86.97% 的用户
内存消耗 : 2.8 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
