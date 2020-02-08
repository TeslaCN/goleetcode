package p338_1

var CountBits = countBits

// Runtime: O(n * sizeof(integer))
func countBits(num int) []int {
	if num < 0 {
		return []int{}
	}
	nums := make([]int, num+1)
	for i := 0; i < len(nums); i++ {
		number := i
		for number > 0 {
			if number&1 == 1 {
				nums[i]++
			}
			number >>= 1
		}
	}
	return nums
}
