package p820_2

var MinimumLengthEncoding = minimumLengthEncoding

type node struct {
	values [26]*node
	end    bool
}

func minimumLengthEncoding(words []string) int {
	if len(words) <= 1 {
		return len(append(words, "")[0]) + 1
	}
	length := 0
	root := &node{}
	for _, word := range words {
		p := root
		create := 0
		isNew := false
		for i := len(word) - 1; i >= 0; i-- {
			pos := word[i] - 0x61
			if p.values[pos] == nil {
				p.values[pos] = &node{}
				isNew = isNew || create == 0 && !p.end
				create++
				p.end = false
			}
			p = p.values[pos]
		}
		p.end = create > 0
		if isNew {
			create = len(word) + 1
		}
		length += create
	}
	return length
}

/*
执行用时 : 32 ms , 在所有 Go 提交中击败了 58.49% 的用户
内存消耗 : 9.3 MB , 在所有 Go 提交中击败了 14.29% 的用户
*/
