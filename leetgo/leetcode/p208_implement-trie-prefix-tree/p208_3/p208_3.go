package p208_3

type Trie struct {
	Nodes [26]*Trie
	End   bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	t := this
	for i := 0; i < len(word); i++ {
		index := word[i] - 0x61
		if trie := t.Nodes[index]; trie == nil {
			trie := Constructor()
			t.Nodes[index] = &trie
		}
		t = t.Nodes[index]
	}
	t.End = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	t := this
	for i := 0; i < len(word); i++ {
		if trie := t.Nodes[word[i]-0x61]; trie == nil {
			return false
		} else {
			t = trie
		}
	}
	return t.End
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	t := this
	for i := 0; i < len(prefix); i++ {
		if trie := t.Nodes[prefix[i]-0x61]; trie == nil {
			return false
		} else {
			t = trie
		}
	}
	return true
}
