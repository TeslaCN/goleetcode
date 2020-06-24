package p238_0

var ProductExceptSelf = productExceptSelf

// Runtime: O(n)
// Space: O(1), ignore res
// Deprecated: Using division
func productExceptSelf(nums []int) []int {
	base := 1
	zero := 0
	for _, num := range nums {
		if num == 0 {
			zero++
		} else {
			base *= num
		}
	}
	res := make([]int, len(nums))
	if zero > 1 {
		return res
	}
	for i := 0; i < len(res); i++ {
		if zero > 0 && nums[i] != 0 {
			res[i] = 0
		} else if nums[i] == 0 {
			res[i] = base
		} else {
			res[i] = base / nums[i]
		}
	}
	return res
}
