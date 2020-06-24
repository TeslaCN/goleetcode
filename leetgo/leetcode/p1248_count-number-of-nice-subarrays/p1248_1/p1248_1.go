package p1248_1

func numberOfSubarrays(nums []int, k int) int {
	n := len(nums)
	if n < 1 || k < 1 || n < k {
		return 0
	}
	result, odd, from, end := 0, 0, 0, 0

	for from < n && end < n {
		for odd < k && end < n {
			if nums[end]&1 == 1 {
				odd++
			}
			if odd < k {
				end++
			}
		}
		for from < n {
			for i := end; i == end && odd == k || i < n && nums[i]&1 == 0; i++ {
				result++
			}
			if nums[from]&1 == 1 {
				from++
				odd--
				end++
				break
			} else {
				from++
			}
		}
	}
	return result
}

/*
执行用时 : 144 ms , 在所有 Go 提交中击败了 84.84% 的用户
内存消耗 : 7 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
