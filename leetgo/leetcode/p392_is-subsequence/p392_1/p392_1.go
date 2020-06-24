package p392_1

var IsSubsequence = isSubsequence

// Runtime: O(n)
// Space: O(1)
func isSubsequence(s string, t string) bool {
	if len(s) > len(t) || (len(s) == len(t) && s != t) {
		return false
	}
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}
