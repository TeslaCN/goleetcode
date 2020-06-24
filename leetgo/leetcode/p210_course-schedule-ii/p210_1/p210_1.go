package p210_1

func findOrder(numCourses int, prerequisites [][]int) []int {
	order := make([]int, 0, numCourses)
	graph, inCounts := make([][]bool, numCourses), make([]int, numCourses)
	for i := 0; i < len(graph); i++ {
		graph[i] = make([]bool, numCourses)
	}
	for _, preReq := range prerequisites {
		graph[preReq[0]][preReq[1]] = true
		inCounts[preReq[0]]++
	}
	for len(order) < numCourses {
		next := -1
		for i, count := range inCounts {
			if count == 0 {
				next = i
				inCounts[i]--
				order = append(order, i)
				break
			}
		}
		if next != -1 {
			for i := 0; i < len(graph); i++ {
				if graph[i][next] {
					graph[i][next] = false
					inCounts[i]--
				}
			}
		} else {
			return []int{}
		}
	}
	return order
}

/*
执行用时 : 32 ms , 在所有 Go 提交中击败了 16.08% 的用户
内存消耗 : 9.1 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
