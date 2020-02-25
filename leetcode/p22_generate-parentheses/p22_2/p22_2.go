package p22_2

var GenerateParenthesis = generateParenthesis

func generateParenthesis(n int) []string {
	return generate(make([]string, 0), make([]byte, 0, n*2), n, n)
}

func generate(result []string, buf []byte, left, right int) []string {
	if left == 0 && right == 0 {
		return append(result, string(buf))
	}
	if left > 0 {
		result = generate(result, append(buf, '('), left-1, right)
	}
	if right > left {
		result = generate(result, append(buf, ')'), left, right-1)
	}
	return result
}
