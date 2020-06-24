package p394_1

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	if !strings.Contains(s, "[") {
		return s
	}
	result := strings.Builder{}
	b := []byte(s)
	times, brackets := 0, 0
	for i := 0; i < len(b); i++ {
		switch {
		case isDigit(b[i]):
			digitFrom := i
			for ; i < len(b) && isDigit(b[i]); i++ {
			}
			times, _ = strconv.Atoi(s[digitFrom:i])
			strFrom, strEnd := i+1, -1
			brackets = 1
			for i = strFrom; i < len(b); i++ {
				switch b[i] {
				case '[':
					brackets++
				case ']':
					brackets--
				default:
				}
				if brackets < 1 {
					strEnd = i
					break
				}
			}
			result.WriteString(strings.Repeat(decodeString(s[strFrom:strEnd]), times))
		default:
			result.WriteByte(b[i])
		}
	}
	return result.String()
}

func isDigit(b byte) bool {
	return 0x30 <= b && b < 0x40
}

/*
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 2 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
