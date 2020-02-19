package p345_1

var ReverseVowels = reverseVowels

func reverseVowels(s string) string {
	b := []byte(s)
	m := map[byte]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
		'A': {},
		'E': {},
		'I': {},
		'O': {},
		'U': {},
	}
	left, right := 0, len(s)-1
	for left < right {
		for left < right {
			_, ok := m[b[left]]
			if ok {
				break
			}
			left++
		}
		for right > left {
			_, ok := m[b[right]]
			if ok {
				break
			}
			right--
		}
		if left < right {
			b[left], b[right] = b[right], b[left]
			left++
			right--
		}
	}
	return string(b)
}
