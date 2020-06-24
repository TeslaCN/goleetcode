package p31_1

// Space: O(1)
// Runtime: O(n)
func nextPermutation(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	from := n - 1
	for ; from > 0; from-- {
		if nums[from-1] < nums[from] {
			break
		}
	}
	if from == 0 {
		// already descending, reverse
		for i := 0; i < n>>1; i++ {
			nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
		}
	} else {
		minDiff, minIndex := nums[from]-nums[from-1], from
		for i := from + 1; i < n && nums[i] > nums[from-1]; i++ {
			if value := nums[i] - nums[from-1]; 0 < value && value <= minDiff {
				minDiff, minIndex = value, i
			}
		}
		nums[from-1], nums[minIndex] = nums[minIndex], nums[from-1]

		// sort.Ints(nums[from:]) // too expensive
		// nums[from:] already descending
		for i := from; i < (from+n)>>1; i++ {
			nums[i], nums[n-1-i+from] = nums[n-1-i+from], nums[i]
		}
	}
}

/*
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 2.5 MB , 在所有 Go 提交中击败了 100.00% 的用户

执行用时 : 4 ms , 在所有 Go 提交中击败了 70.05% 的用户
内存消耗 : 2.5 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
