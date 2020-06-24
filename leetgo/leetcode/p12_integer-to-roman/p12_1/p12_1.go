package p12_1

var IntToRoman = intToRoman

func intToRoman(num int) string {
	if num < 1 {
		return ""
	}
	options := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	i := 0
	result := ""
	for num > 0 {
		if num < options[i] {
			i++
		} else {
			num -= options[i]
			result += romans[i]
		}
	}
	return result
}
