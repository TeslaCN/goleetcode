package p208_2

type Trie struct {
	Nodes map[byte]*Trie
	End   bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{make(map[byte]*Trie), false}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	t := this
	for i := 0; i < len(word); i++ {
		if _, exists := t.Nodes[word[i]]; !exists {
			trie := Constructor()
			t.Nodes[word[i]] = &trie
		}
		t = t.Nodes[word[i]]
	}
	t.End = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	t := this
	for i := 0; i < len(word); i++ {
		if next, exists := t.Nodes[word[i]]; !exists {
			return false
		} else {
			t = next
		}
	}
	return t.End
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	t := this
	for i := 0; i < len(prefix); i++ {
		if next, exists := t.Nodes[prefix[i]]; !exists {
			return false
		} else {
			t = next
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
