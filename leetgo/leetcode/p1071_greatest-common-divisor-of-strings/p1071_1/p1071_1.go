package p1071_1

func gcdOfStrings(str1 string, str2 string) string {
	l1, l2 := len(str1), len(str2)
	if l1 < l2 {
		l1, l2 = l2, l1
	}
	for l2 != 0 {
		l1, l2 = l2, l1%l2
	}
	maxLen := l1
	ok := true
	for i := 0; ok && i < len(str1); i++ {
		if str1[i] != str1[i%maxLen] {
			ok = false
		}
	}
	if !ok {
		return ""
	}
	ok = true
	for i := 0; ok && i < len(str2); i++ {
		if str2[i] != str2[i%maxLen] {
			ok = false
		}
	}
	if !ok {
		return ""
	}
	if str1[:maxLen] == str2[:maxLen] {
		return str1[:maxLen]
	} else {
		return ""
	}
}
