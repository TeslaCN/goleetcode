package p152_1

// Runtime: O(n)
// Space: O(1)
func maxProduct(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	maxValue, from, to := nums[0], 0, 0
	for to < n || from < n {
		product := 1
		for to < n && nums[to] != 0 {
			product *= nums[to]
			if product > maxValue {
				maxValue = product
			}
			to++
		}
		for from < to-1 {
			product /= nums[from]
			if product > maxValue {
				maxValue = product
			}
			from++
		}
		if to < n && nums[to] > maxValue {
			maxValue = nums[to]
		}
		for ; to < n && nums[to] == 0; to++ {
		}
		from = to
	}

	return maxValue
}

/*
执行用时 : 4 ms , 在所有 Go 提交中击败了 88.25% 的用户
内存消耗 : 2.7 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
