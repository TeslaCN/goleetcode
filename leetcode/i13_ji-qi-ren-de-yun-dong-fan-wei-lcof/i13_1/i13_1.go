package i13_1

const (
	WALKED  int8 = 1
	UNKNOWN int8 = 0
	BLOCKED int8 = -1
)

func movingCount(m int, n int, k int) int {
	if k == 0 {
		return 1
	}
	if m < 1 || n < 1 {
		return 0
	}
	cache := make([][]int8, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int8, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if numSum(i)+numSum(j) > k {
				cache[i][j] = BLOCKED
			}
		}
	}

	return check(cache, 0, 0)
}

func check(cache [][]int8, x, y int) int {
	if x < 0 || y < 0 || x >= len(cache) || y >= len(cache[x]) || cache[x][y] != UNKNOWN {
		return 0
	}
	cache[x][y] = WALKED
	return 1 + check(cache, x-1, y) + check(cache, x, y-1) + check(cache, x+1, y) + check(cache, x, y+1)
}

func numSum(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

/*
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 2.5 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
