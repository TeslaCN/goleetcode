package p22_3

var GenerateParenthesis = generateParenthesis

func generateParenthesis(n int) []string {
	return gen(make([]string, 0), make([]byte, 0, n*2), n, n)
}

func gen(result []string, cache []byte, left, right int) []string {
	if left == 0 && right == 0 {
		return append(result, string(cache))
	}
	if left > 0 {
		result = gen(result, append(cache, '('), left-1, right)
	}
	if right > left {
		result = gen(result, append(cache, ')'), left, right-1)
	}
	return result
}
