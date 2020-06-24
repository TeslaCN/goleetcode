package p208_implement_trie_prefix_tree

import (
	"github.com/TeslaCN/goleetcode/leetcode/p208_implement-trie-prefix-tree/p208_1"
	"github.com/TeslaCN/goleetcode/leetcode/p208_implement-trie-prefix-tree/p208_2"
	"github.com/TeslaCN/goleetcode/leetcode/p208_implement-trie-prefix-tree/p208_3"
	"testing"
)

func TestTrie_1(t *testing.T) {
	trie := p208_1.Constructor()
	t.Log(trie.Search("apple") == false)
	t.Log(trie.StartsWith("apple") == false)
	trie.Insert("apple")
	t.Log(trie.Search("apple") == true)
	t.Log(trie.StartsWith("apple") == true)
	t.Log(trie.Search("app") == false)
	t.Log(trie.StartsWith("app") == true)
}

func TestTrie_2(t *testing.T) {
	trie := p208_2.Constructor()
	t.Log(trie.Search("apple") == false)
	t.Log(trie.StartsWith("apple") == false)
	trie.Insert("apple")
	t.Log(trie.Search("apple") == true)
	t.Log(trie.StartsWith("apple") == true)
	t.Log(trie.Search("app") == false)
	t.Log(trie.StartsWith("app") == true)
}

func TestTrie_3(t *testing.T) {
	trie := p208_3.Constructor()
	t.Log(trie.Search("apple") == false)
	t.Log(trie.StartsWith("apple") == false)
	trie.Insert("apple")
	t.Log(trie.Search("apple") == true)
	t.Log(trie.StartsWith("apple") == true)
	t.Log(trie.Search("app") == false)
	t.Log(trie.StartsWith("app") == true)
}
