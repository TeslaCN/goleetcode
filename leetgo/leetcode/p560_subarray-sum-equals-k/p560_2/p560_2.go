package p560_2

var SubarraySum = subarraySum

func subarraySum(nums []int, k int) int {
	n, hits := len(nums), 0
	if n < 1 {
		return 0
	}
	cache, sum := map[int]int{0: 1}, 0
	for _, num := range nums {
		sum += num
		hits += cache[sum-k]
		cache[sum]++
	}
	return hits
}

/*
执行用时 : 20 ms , 在所有 Go 提交中击败了 85.23% 的用户
内存消耗 : 6.3 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
