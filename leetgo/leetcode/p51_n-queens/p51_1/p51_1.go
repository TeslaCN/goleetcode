package p51_1

import (
	"bytes"
)

func solveNQueens(n int) [][]string {
	switch {
	case n <= 0:
		return [][]string{}
	case n == 1:
		return [][]string{{"Q"}}
	}
	result := make([][]string, 0)
	return backTrack(make([]int, n), 0, result)
}

func backTrack(pos []int, level int, result [][]string) [][]string {
	if level == len(pos) {
		solution := make([]string, len(pos))
		for i := 0; i < len(pos); i++ {
			b := bytes.Repeat([]byte{'.'}, len(pos))
			b[pos[i]] = 'Q'
			solution[i] = string(b)
		}
		result = append(result, solution)
	} else {
		for i := 0; i < len(pos); i++ {
			pos[level] = i
			if place(pos, level, i) {
				result = backTrack(pos, level+1, result)
			}
		}
	}
	return result
}

func place(pos []int, level, target int) bool {
	for i := 0; i < level; i++ {
		if pos[i] == target || pos[i]-i == target-level || pos[i]+i == target+level {
			return false
		}
	}
	return true
}
