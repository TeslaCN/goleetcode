package p524_1

var FindLongestWord = findLongestWord

// 双指针
// Space: O(len(max))
// Runtime: O(x*len(d))
func findLongestWord(s string, d []string) string {
	max := ""
	for _, target := range d {
		if len(target) > len(s) || len(target) < len(max) {
			continue
		}
		ps, pt := 0, 0
		for ps < len(s) && pt < len(target) {
			if target[pt] == s[ps] {
				pt++
			}
			ps++
		}
		if pt == len(target) && (len(target) > len(max) || target < max) {
			max = target
		}
	}
	return max
}
