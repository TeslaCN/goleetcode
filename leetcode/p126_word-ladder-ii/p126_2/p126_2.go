package p126_2

import "math"

var FindLadders = findLadders

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordIds, idWords := map[string]int{}, make([]string, len(wordList))
	for i, word := range wordList {
		wordIds[word], idWords[i] = i, word
	}
	if _, exists := wordIds[endWord]; !exists {
		return [][]string{}
	}
	if _, exists := wordIds[beginWord]; !exists {
		wordIds[beginWord], idWords = len(wordIds), append(idWords, beginWord)
		wordList = append(wordList, beginWord)
	}
	neighbors := make([][]int, len(wordList))
	for i := 0; i < len(wordList)-1; i++ {
		for j := i + 1; j < len(wordList); j++ {
			if near(wordList[i], wordList[j]) {
				neighbors[i] = append(neighbors[i], j)
				neighbors[j] = append(neighbors[j], i)
			}
		}
	}

	cost, endId := make([]int, len(wordList)), wordIds[endWord]
	for i := 0; i < len(cost); i++ {
		cost[i] = math.MaxInt32 - 1
	}
	cost[wordIds[beginWord]] = 0

	queue, answers := [][]int{{wordIds[beginWord]}}, make([][]string, 0)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		lastVisited := cur[len(cur)-1]
		if lastVisited == endId {
			anAns := make([]string, len(cur))
			for i := 0; i < len(anAns); i++ {
				anAns[i] = idWords[cur[i]]
			}
			answers = append(answers, anAns)
		} else {
			for _, nextId := range neighbors[lastVisited] {
				if cost[nextId] >= cost[lastVisited]+1 {
					cost[nextId] = cost[lastVisited] + 1
					copied := make([]int, len(cur))
					copy(copied, cur)
					copied = append(copied, nextId)
					queue = append(queue, copied)
				}
			}
		}
	}
	return answers
}

func near(s1, s2 string) bool {
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return s1[i+1:] == s2[i+1:]
		}
	}
	return true
}

/*
执行用时 : 244 ms , 在所有 Go 提交中击败了 46.55% 的用户
内存消耗 : 7.5 MB , 在所有 Go 提交中击败了 33.33% 的用户
*/
