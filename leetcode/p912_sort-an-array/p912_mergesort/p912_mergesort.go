package p912_mergesort

var SortArray = sortArray

func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	cache := make([]int, len(nums))
	for l := 2; ; l <<= 1 {
		for i := 0; i < len(nums); i += l {
			ms(nums, i, min(i+(l>>1), len(nums)), min(i+l, len(nums)), cache)
		}
		if l >= len(nums) {
			break
		}
	}
	return nums
}

func ms(nums []int, start, mid, end int, cache []int) {
	if end-start <= 1 {
		return
	}
	cur, a, b := 0, start, mid
	for a < mid && b < end {
		if nums[a] < nums[b] {
			cache[cur] = nums[a]
			a++
		} else {
			cache[cur] = nums[b]
			b++
		}
		cur++
	}
	for a < mid {
		cache[cur] = nums[a]
		cur++
		a++
	}
	for b < end {
		cache[cur] = nums[b]
		cur++
		b++
	}
	copy(nums[start:end], cache[:cur])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
执行用时 : 24 ms , 在所有 Go 提交中击败了 90.65% 的用户
内存消耗 : 6.6 MB , 在所有 Go 提交中击败了 5.71% 的用户
*/
