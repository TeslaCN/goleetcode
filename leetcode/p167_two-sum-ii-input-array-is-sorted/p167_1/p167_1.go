package p167_1

var TwoSum = twoSum

// 双指针
// Runtime: O(n)
// Space: O(1)
func twoSum(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return []int{0, 0}
	}
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		}
		if sum > target {
			right--
		} else {
			left++
		}
	}
	return []int{0, 0}
}
