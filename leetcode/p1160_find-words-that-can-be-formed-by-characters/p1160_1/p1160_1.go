package p1160_1

func countCharacters(words []string, chars string) int {
	var c [26]int
	for i := 0; i < len(chars); i++ {
		c[chars[i]-'a']++
	}
	length := 0
	for _, word := range words {
		tmp := c
		learned := true
		for i := 0; i < len(word); i++ {
			pos := word[i] - 'a'
			tmp[pos]--
			if tmp[pos] < 0 {
				learned = false
				break
			}
		}
		if learned {
			length += len(word)
		}
	}
	return length
}
