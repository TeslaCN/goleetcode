package p238_1

var ProductExceptSelf = productExceptSelf

func productExceptSelf(nums []int) []int {
	n := len(nums)
	fromLeft, fromRight := make([]int, n), make([]int, n)
	fromLeft[0], fromRight[n-1] = nums[0], nums[n-1]
	for i := 1; i < n; i++ {
		fromLeft[i] = fromLeft[i-1] * nums[i]
		fromRight[n-1-i] = fromRight[n-i] * nums[n-1-i]
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		switch i {
		case 0:
			res[i] = fromRight[i+1]
		case n - 1:
			res[i] = fromLeft[i-1]
		default:
			res[i] = fromLeft[i-1] * fromRight[i+1]
		}
	}
	return res
}
