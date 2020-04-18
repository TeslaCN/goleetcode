package p79_1

func exist(board [][]byte, word string) bool {
	if word == "" {
		return true
	}
	if len(board) < 1 || len(board[0]) < 1 {
		return false
	}
	bytes := []byte(word)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if dfs(board, i, j, bytes) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, i, j int, bytes []byte) bool {
	if len(bytes) < 1 {
		return true
	}
	if i >= 0 && i < len(board) && j >= 0 && j < len(board[i]) && bytes[0] == board[i][j] {
		temp := board[i][j]
		board[i][j] = 0
		newBytes := bytes[1:]
		result := dfs(board, i-1, j, newBytes) || dfs(board, i+1, j, newBytes) || dfs(board, i, j-1, newBytes) || dfs(board, i, j+1, newBytes)
		board[i][j] = temp
		return result
	}
	return false
}

/*
执行用时 : 4 ms , 在所有 Go 提交中击败了 96.15% 的用户
内存消耗 : 3.5 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
