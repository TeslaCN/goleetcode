package p8_1

import "math"

var MyAtoi = myAtoi

func myAtoi(str string) int {
	start := 0
	for ; start < len(str); start++ {
		if str[start] != ' ' {
			break
		}
	}
	str = str[start:]

	started := false
	sign := 1
	result := 0
loop:
	for i := 0; i < len(str); i++ {
		switch {
		case str[i] == '-':
			if !started {
				sign = -1
				started = true
			} else {
				break loop
			}
		case str[i] == '+':
			if !started {
				sign = 1
				started = true
			} else {
				break loop
			}
		case 0x30 <= str[i] && str[i] < 0x40:
			if !started {
				started = true
			}
			if result < math.MaxInt32/10 {
				result *= 10
				result += int(str[i]) - 0x30
			} else if result > math.MaxInt32/10 {
				if sign == 1 {
					return math.MaxInt32
				} else {
					return math.MinInt32
				}
			} else {
				toAdd := int(str[i]) - 0x30
				if sign == 1 && toAdd >= 7 {
					return math.MaxInt32
				}
				if sign == -1 && toAdd >= 8 {
					return math.MinInt32
				}
				result *= 10
				result += int(str[i]) - 0x30
			}
		default:
			break loop
		}
	}
	return result * sign
}

/*
执行用时 : 4 ms , 在所有 Go 提交中击败了 55.01% 的用户
内存消耗 : 2.3 MB , 在所有 Go 提交中击败了 84.85% 的用户
*/
