package p524_longest_word_in_dictionary_through_deleting

import (
	"github.com/TeslaCN/goleetcode/leetcode/p524_longest-word-in-dictionary-through-deleting/p524_1"
	"testing"
)

type args struct {
	s string
	d []string
}

var tests = []struct {
	name string
	args args
	want string
}{
	{"sample-0", args{"abpcplea", []string{"ale", "apple", "monkey", "plea"}}, "apple"},
	{"sample-1", args{"abpcplea", []string{"a", "b", "c"}}, "a"},
}

func Test_findLongestWord(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p524_1.FindLongestWord(tt.args.s, tt.args.d); got != tt.want {
				t.Errorf("findLongestWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
