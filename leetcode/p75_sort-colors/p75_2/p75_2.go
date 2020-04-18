package p75_2

var SortColors = sortColors

func sortColors(nums []int) {
	if len(nums) < 2 {
		return
	}
	left, right := 0, len(nums)-1
	for i := left; i <= right; {
		switch nums[i] {
		case 0:
			nums[i], nums[left] = nums[left], nums[i]
			left++
		case 2:
			nums[i], nums[right] = nums[right], nums[i]
			right--
		default:
			i++
		}
		if i <= left {
			i = left
		}
	}
}

/*
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 2.1 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
