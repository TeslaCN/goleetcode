package p22_1

var GenerateParenthesis = generateParenthesis

func generateParenthesis(n int) []string {
	switch n {
	case 0:
		return []string{}
	case 1:
		return []string{"()"}
	}
	parenthesis := generateParenthesis(n - 1)
	resultMap := make(map[string]struct{})
	for _, p := range parenthesis {
		for i := 0; i <= len(p); i++ {
			s := p[:i] + "()" + p[i:]
			resultMap[s] = struct{}{}
		}
	}
	res := make([]string, 0)
	for key := range resultMap {
		res = append(res, key)
	}
	return res
}
