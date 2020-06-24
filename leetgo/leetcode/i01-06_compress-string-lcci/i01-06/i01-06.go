package i01_06

import (
	"strconv"
	"strings"
)

func compressString(S string) string {
	b := strings.Builder{}
	pre, end := 0, 0
	for end < len(S) && b.Len() < len(S) {
		c := S[end]
		for end < len(S) && S[end] == c {
			end++
		}
		b.WriteByte(c)
		b.WriteString(strconv.Itoa(end - pre))
		pre = end
	}
	if end < len(S) || b.Len() >= len(S) {
		return S
	} else {
		return b.String()
	}
}
