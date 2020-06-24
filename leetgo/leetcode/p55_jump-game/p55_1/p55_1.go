package p55_1

var CanJump = canJump

func canJump(nums []int) bool {
	n := len(nums)
	if n <= 1 || nums[0] >= n-1 {
		return true
	} else if nums[0] == 0 {
		return false
	}
	end := nums[0]
	for i := 0; i <= end; i++ {
		far := i + nums[i]
		if far >= n-1 {
			return true
		}
		if far > end {
			end = far
		}
	}
	return false
}

/*
执行用时 : 8 ms , 在所有 Go 提交中击败了 97.49% 的用户
内存消耗 : 4.2 MB , 在所有 Go 提交中击败了 100.00% 的用户

执行用时 : 16 ms , 在所有 Go 提交中击败了 35.50% 的用户
内存消耗 : 4.2 MB , 在所有 Go 提交中击败了 46.15% 的用户
*/
