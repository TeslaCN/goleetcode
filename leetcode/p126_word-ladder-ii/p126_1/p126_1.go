package p126_1

// Deprecated: TLE
var FindLadders = findLadders

// Deprecated: TLE
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	words := map[string]bool{beginWord: true}
	for _, w := range wordList {
		words[w] = true
	}
	if !words[endWord] || beginWord == endWord {
		return [][]string{}
	}
	neibors := map[string]map[string]bool{}
	for word := range words {
		for w := range words {
			if editDistanceEqualsLen(word, w) == 1 {
				if len(neibors[word]) == 0 {
					neibors[word] = map[string]bool{}
				}
				neibors[word][w] = true
			}
		}
	}
	words[beginWord] = false
	ans, min := map[int][][]string{}, -1
	backTrack(ans, append(make([]string, 0, len(words)+1), beginWord), words, neibors, endWord, &min)
	if len(ans) == 0 {
		return [][]string{}
	}
	return ans[min]
}

func backTrack(ans map[int][][]string, queue []string, words map[string]bool, neibors map[string]map[string]bool, endWord string, min *int) {
	lastWord := queue[len(queue)-1]
	for next, valid := range neibors[lastWord] {
		if valid && words[next] {
			queue = append(queue, next)
			if next == endWord {
				if *min < 0 || len(queue) < *min {
					*min = len(queue)
				}
				copied := make([]string, len(queue))
				copy(copied, queue)
				ans[len(queue)] = append(ans[len(queue)], copied)
				return
			} else if *min < 0 || len(queue) < *min {
				words[next] = false
				backTrack(ans, queue, words, neibors, endWord, min)
				words[next] = true
			}
			queue = queue[:len(queue)-1]
		}
	}
}

func editDistanceEqualsLen(s1, s2 string) int {
	distance := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			distance++
		}
	}
	return distance
}
