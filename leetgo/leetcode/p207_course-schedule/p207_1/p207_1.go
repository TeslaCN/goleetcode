package p207_1

var CanFinish = canFinish

func canFinish(numCourses int, prerequisites [][]int) bool {
	matrix, reqRemains, remain := make([][]bool, numCourses), make([]int, numCourses), numCourses
	for i := 0; i < numCourses; i++ {
		matrix[i] = make([]bool, numCourses)
	}
	for _, prerequisite := range prerequisites {
		matrix[prerequisite[0]][prerequisite[1]] = true
		reqRemains[prerequisite[0]]++
	}
	for remain > 0 {
		chosen := -1
		for i := 0; i < numCourses; i++ {
			if reqRemains[i] == 0 {
				reqRemains[i], chosen = -1, i
				remain--
				break
			}
		}
		if chosen == -1 {
			return false
		}
		for i := 0; i < numCourses; i++ {
			if matrix[i][chosen] {
				matrix[i][chosen] = false
				reqRemains[i]--
			}
		}
	}
	return true
}

/*
执行用时 : 36 ms , 在所有 Go 提交中击败了 15.93% 的用户
内存消耗 : 9.1 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
