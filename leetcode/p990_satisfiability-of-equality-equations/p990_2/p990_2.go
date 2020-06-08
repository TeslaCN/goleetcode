package p990_2

var EquationsPossible = equationsPossible

func equationsPossible(equations []string) bool {
	parents := [26]int{}
	for i := 0; i < len(parents); i++ {
		parents[i] = i
	}
	for _, expr := range equations {
		if expr[1] == '=' {
			index1, index2 := int(expr[0]-'a'), int(expr[3]-'a')
			union(&parents, index1, index2)
		}
	}
	for _, expr := range equations {
		if expr[1] == '!' {
			index1, index2 := int(expr[0]-'a'), int(expr[3]-'a')
			if find(&parents, index1) == find(&parents, index2) {
				return false
			}
		}
	}
	return true
}

func union(parents *[26]int, index1, index2 int) {
	parents[find(parents, index2)] = find(parents, index1)
}

func find(parents *[26]int, index int) int {
	for parents[index] != index {
		// optimize
		parents[index] = parents[parents[index]]
		index = parents[index]
	}
	return index
}

/*
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 2.8 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
