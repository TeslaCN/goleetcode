package p208_1

// Recursive
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
	if len(word) < 1 {
		this.End = true
		return
	}
	if _, exists := this.Nodes[word[0]]; !exists {
		trie := Constructor()
		this.Nodes[word[0]] = &trie
	}
	this.Nodes[word[0]].Insert(word[1:])
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if len(word) < 1 {
		return this.End
	}
	next, exists := this.Nodes[word[0]]
	return exists && next.Search(word[1:])
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) < 1 {
		return true
	}
	next, exists := this.Nodes[prefix[0]]
	return exists && next.StartsWith(prefix[1:])
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
