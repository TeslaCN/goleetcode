package p283_1

func MoveZeroes(nums []int) {
	if nums == nil || len(nums) < 1 {
		return
	}
	i, next := 0, 0
	for next < len(nums) {
		if nums[next] == 0 {
			next++
			continue
		}
		nums[i] = nums[next]
		i++
		next++
	}
	for ; i < len(nums); i++ {
		nums[i] = 0
	}
}
