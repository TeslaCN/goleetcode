package p621_1

import "sort"

func leastInterval(tasks []byte, n int) int {
	l := len(tasks)
	if n == 0 {
		return l
	}
	time := 0
	taskCount := make([]int, 26)
	for _, task := range tasks {
		taskCount[task-'A']++
	}
	sort.Ints(taskCount)
	for l > 0 {
		remain := n + 1
		for i := len(taskCount) - 1; l > 0 && remain > 0 && i >= 0; i-- {
			if taskCount[i] > 0 {
				taskCount[i]--
				remain--
				l--
			}
		}
		if l > 0 {
			time += n + 1
			sort.Ints(taskCount)
		} else {
			time += n + 1 - remain
		}
	}
	return time
}

/*
执行用时 : 16 ms , 在所有 Go 提交中击败了 46.48% 的用户
内存消耗 : 6 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
