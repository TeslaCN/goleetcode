package p39_1

func combinationSum(candidates []int, target int) [][]int {
	return combination(make([][]int, 0), candidates, make([]int, 0, len(candidates)), target)
}

func combination(result [][]int, candidates, buf []int, target int) [][]int {
	for _, candidate := range candidates {
		if candidate > target || (len(buf) > 0 && candidate < buf[len(buf)-1]) {
			continue
		} else if candidate == target {
			comb := make([]int, len(buf))
			copy(comb, buf)
			result = append(result, append(comb, candidate))
		} else {
			result = combination(result, candidates, append(buf, candidate), target-candidate)
		}
	}
	return result
}
