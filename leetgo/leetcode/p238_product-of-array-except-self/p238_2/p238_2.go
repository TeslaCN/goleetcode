package p238_2

var ProductExceptSelf = productExceptSelf

func productExceptSelf(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}
	if n == 1 {
		return nums
	}
	ans := make([]int, len(nums))
	ans[0] = nums[0]
	for i := 1; i < len(ans)-1; i++ {
		ans[i] = ans[i-1] * nums[i]
	}
	ans[n-1] = ans[n-2]
	product := nums[n-1]
	for i := n - 2; i > 0; i-- {
		ans[i] = product * ans[i-1]
		product *= nums[i]
	}
	ans[0] = product
	return ans
}

/*
执行用时 : 12 ms , 在所有 Go 提交中击败了 93.26% 的用户
内存消耗 : 6.3 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
