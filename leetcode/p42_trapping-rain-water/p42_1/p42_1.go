package p42_1

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	sum, left, right := 0, 0, 0
	for left < len(height) && right < len(height) {
		for ; left+1 < len(height) && (height[left] <= 0 || height[left] <= height[left+1]); left++ {
		}
		right = left + 2
		if right >= len(height) {
			break
		}
		for i := left + 3; i < len(height); i++ {
			if height[right] >= height[left] {
				break
			}
			if height[i] >= height[right] {
				right = i
			}
		}
		min := height[left]
		if height[right] < min {
			min = height[right]
		}
		for i := left + 1; i < right; i++ {
			if value := min - height[i]; value > 0 {
				sum += value
			}
		}
		left = right
	}
	return sum
}

/*
执行用时 : 4 ms , 在所有 Go 提交中击败了 86.97% 的用户
内存消耗 : 2.8 MB , 在所有 Go 提交中击败了 74.07% 的用户
*/
