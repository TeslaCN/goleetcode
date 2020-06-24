package p680_1

var ValidPalindrome = validPalindrome

// 双指针
// Runtime: O(n)
// Space: O(1)
func validPalindrome(s string) bool {
	return validPalindromeDeletable(s, true)
}

func validPalindromeDeletable(s string, deletable bool) bool {
	if len(s) < 2 {
		return true
	}
	left, right := 0, len(s)-1
	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			return deletable && (validPalindromeDeletable(s[left+1:right+1], false) || validPalindromeDeletable(s[left:right], false))
		}
	}
	return true
}
