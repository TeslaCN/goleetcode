package p1108_1

import "strings"

var DefangIPaddr = defangIPaddr

func defangIPaddr(address string) string {
	sb := strings.Builder{}
	for i := 0; i < len(address); i++ {
		if address[i] == '.' {
			sb.WriteString("[.]")
		} else {
			sb.WriteByte(address[i])
		}
	}
	return sb.String()
}
