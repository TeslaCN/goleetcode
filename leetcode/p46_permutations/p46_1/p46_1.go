package p46_1

var Permute = permute

func permute(nums []int) [][]int {
	if len(nums) < 1 {
		return [][]int{}
	}
	var sets [][]int

	var p func(nums []int, from int)
	p = func(nums []int, from int) {
		if from+1 >= len(nums) {
			copied := make([]int, len(nums))
			copy(copied, nums)
			sets = append(sets, copied)
			return
		}
		for i := from; i < len(nums); i++ {
			swap(&nums[from], &nums[i])
			p(nums, from+1)
			swap(&nums[from], &nums[i])
		}
	}
	p(nums, 0)
	return sets
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

//func p(nums []int, from int) {
//	if from+1 == len(nums) {
//		copied := make([]int, len(nums))
//		copy(copied, nums)
//		sets = append(sets, copied)
//	}
//	for i := from; i < len(nums); i++ {
//		swap(&nums[from], &nums[i])
//		p(nums, i)
//		swap(&nums[from], &nums[i])
//	}
//}
