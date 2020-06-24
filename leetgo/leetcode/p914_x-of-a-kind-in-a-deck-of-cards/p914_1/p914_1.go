package p914_1

func hasGroupsSizeX(deck []int) bool {
	if len(deck) <= 1 {
		return false
	}
	m := make(map[int]int)
	for _, value := range deck {
		m[value]++
	}
	if len(m) == 1 {
		return true
	}
	values := make([]int, len(m))
	i := 0
	for _, count := range m {
		values[i] = count
		i++
	}
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			a, b := values[i], values[j]
			if a < b {
				a, b = b, a
			}
			for b != 0 {
				a, b = b, a%b
			}
			if a < 2 {
				return false
			}
		}
	}
	return true
}
