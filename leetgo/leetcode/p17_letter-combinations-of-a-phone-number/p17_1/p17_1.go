package p17_1

import "strings"

var LetterCombinations = letterCombinations

var m = make(map[string][]string)

func letterCombinations(digits string) []string {
	if strings.TrimSpace(digits) == "" {
		return []string{}
	}
	m["2"] = []string{"a", "b", "c"}
	m["3"] = []string{"d", "e", "f"}
	m["4"] = []string{"g", "h", "i"}
	m["5"] = []string{"j", "k", "l"}
	m["6"] = []string{"m", "n", "o"}
	m["7"] = []string{"p", "q", "r", "s"}
	m["8"] = []string{"t", "u", "v"}
	m["9"] = []string{"w", "x", "y", "z"}
	return combination(digits, []string{})
}

func combination(digits string, combinations []string) []string {
	if combinations == nil {
		return nil
	}
	if strings.TrimSpace(digits) == "" {
		return combinations
	}
	d := digits[len(digits)-1:]
	letters, ok := m[d]
	if ok {
		beforeCount := len(combinations)
		if beforeCount > 0 {

			for i := 0; i < len(letters)-1; i++ {
				for j := 0; j < beforeCount; j++ {
					combinations = append(combinations, combinations[j])
				}
			}
			for i := 0; i < len(letters); i++ {
				for j := i * beforeCount; j < (i+1)*beforeCount; j++ {
					combinations[j] = letters[i] + combinations[j]
				}
			}
		} else {
			for _, letter := range letters {
				combinations = append(combinations, letter)
			}
		}
	}
	return combination(digits[:len(digits)-1], combinations)
}
