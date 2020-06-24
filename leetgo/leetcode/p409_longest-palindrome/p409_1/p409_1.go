package p409_1

func longestPalindrome(s string) int {
	var m [52]int
	length := 0
	for i := 0; i < len(s); i++ {
		pos := index(s[i])
		m[pos]++
		if m[pos] == 2 {
			length += 2
			m[pos] -= 2
		}
	}
	if len(s) > length {
		length++
	}
	return length
}

func index(c byte) int {
	if c >= 0x61 {
		return int(c - 0x61 + 26)
	} else {
		return int(c - 0x41)
	}
}
