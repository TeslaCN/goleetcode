package p945_2

var MinIncrementForUnique = minIncrementForUnique

func minIncrementForUnique(A []int) int {
	add := 0
	if len(A) <= 1 {
		return add
	}
	slot := make([]int, 40000)
	for _, v := range A {
		slot[v]++
	}
	ta, empty := 0, 1
	for ta < 40000 {
		for slot[ta] > 1 {
			if empty < ta {
				empty = ta + 1
			}
			for empty < len(slot) && slot[empty] > 0 {
				empty++
			}
			if empty < len(slot) {
				slot[empty]++
			} else {
				slot = append(slot, 1)
			}
			slot[ta]--
			add += empty - ta
		}
		ta++
	}
	return add
}

/*
执行用时 : 48 ms , 在所有 Go 提交中击败了 96.15% 的用户
内存消耗 : 7 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
