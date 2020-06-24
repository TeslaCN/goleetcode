package p207_2

var CanFinish = canFinish

func canFinish(numCourses int, prerequisites [][]int) bool {
	matrix, reqRemains, queue, remain := make([][]bool, numCourses), make([]int, numCourses), make([]int, 0, numCourses), numCourses
	for i := 0; i < numCourses; i++ {
		matrix[i] = make([]bool, numCourses)
	}
	for _, prerequisite := range prerequisites {
		matrix[prerequisite[0]][prerequisite[1]] = true
		reqRemains[prerequisite[0]]++
	}
	for i := 0; i < numCourses; i++ {
		if reqRemains[i] == 0 {
			queue = append(queue, i)
			reqRemains[i] = -1
			remain--
		}
	}
	for len(queue) > 0 && remain > 0 {
		chosen := queue[0]
		queue = queue[1:]
		for i := 0; i < numCourses; i++ {
			if matrix[i][chosen] {
				reqRemains[i]--
			}
			if reqRemains[i] == 0 {
				queue = append(queue, i)
				reqRemains[i] = -1
				remain--
			}
		}
	}
	return remain == 0
}

/*
执行用时 : 36 ms , 在所有 Go 提交中击败了 15.93% 的用户
内存消耗 : 9.2 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
