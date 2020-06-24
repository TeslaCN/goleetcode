package p128_1

var LongestConsecutive = longestConsecutive

type Node struct {
	val  int
	pre  *Node
	next *Node
}

// Runtime: O(n)
// Space: O(n)
func longestConsecutive(nums []int) int {
	cache := map[int]*Node{}
	for _, num := range nums {
		if _, exists := cache[num]; !exists {
			node := &Node{val: num}
			if cache[num+1] != nil {
				node.next = cache[num+1]
				cache[num+1].pre = node
			}
			if cache[num-1] != nil {
				node.pre = cache[num-1]
				cache[num-1].next = node
			}
			cache[num] = node
		}
	}
	if len(cache) == 0 {
		return 0
	}
	maxLen := 1
	for _, head := range cache {
		if head.pre == nil && head.next != nil {
			len := 0
			for p := head; p != nil; p = p.next {
				len++
			}
			if len > maxLen {
				maxLen = len
			}
		}
	}
	return maxLen
}

/*
执行用时 : 8 ms , 在所有 Go 提交中击败了 68.84% 的用户
内存消耗 : 4.3 MB , 在所有 Go 提交中击败了 50.00% 的用户
*/
