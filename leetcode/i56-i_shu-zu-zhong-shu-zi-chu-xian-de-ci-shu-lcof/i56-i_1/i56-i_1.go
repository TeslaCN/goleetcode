package i56_i_1

func singleNumbers(nums []int) []int {
	value := 0
	for _, num := range nums {
		value ^= num
	}
	partition := value & -value
	result := []int{0, 0}
	for _, num := range nums {
		if num&partition == partition {
			result[0] ^= num
		} else {
			result[1] ^= num
		}
	}
	return result
}

/*
执行用时 : 16 ms , 在所有 Go 提交中击败了 99.55% 的用户
内存消耗 : 6 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
