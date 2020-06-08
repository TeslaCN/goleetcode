package p990_1

var EquationsPossible = equationsPossible

// BFS
func equationsPossible(equations []string) bool {
	var equals ['z' + 1][]byte
	for _, expr := range equations {
		if expr[1] == '=' {
			var1, var2 := expr[0], expr[3]
			if var1 != var2 {
				equals[var1], equals[var2] = append(equals[var1], var2), append(equals[var2], var1)
			}
		}
	}
	for _, expr := range equations {
		if expr[1] == '!' {
			var1, var2 := expr[0], expr[3]
			if var1 == var2 || reachable(equals, var1, var2) {
				return false
			}
		}
	}
	return true
}

func reachable(equals ['z' + 1][]byte, from, to byte) bool {
	queue, visited := []byte{from}, map[byte]bool{}
	for len(queue) > 0 {
		varFrom := queue[0]
		queue, visited[varFrom] = queue[1:], true
		for _, next := range equals[varFrom] {
			if next == to {
				return true
			}
			if !visited[next] {
				queue = append(queue, next)
			}
		}
	}
	return false
}

/*
执行用时 : 4 ms , 在所有 Go 提交中击败了 76.19% 的用户
内存消耗 : 2.9 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
