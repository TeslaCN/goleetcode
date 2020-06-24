package p983_2

var MincostTickets = mincostTickets

var exp = [3]int{1, 7, 30}

// Space: O(n)
// Runtime: O(n)
func mincostTickets(days []int, costs []int) int {
	cache := make([]int, len(days))
	cache[0] = min(costs[0], min(costs[1], costs[2]))
	for i := 1; i < len(days); i++ {
		for c, cost := range costs {
			if c == 0 {
				cache[i] = cache[i-1] + cost
				continue
			}
			cover := i
			for ; cover > 0 && days[cover-1] > days[i]-exp[c]; cover-- {
			}
			if cover == 0 {
				cache[i] = min(cache[i], cost)
			} else {
				cache[i] = min(cache[i], cost+cache[cover-1])
			}
		}
	}
	return cache[len(cache)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 2.3 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
