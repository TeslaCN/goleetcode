package p139_word_break

import (
	"github.com/TeslaCN/goleetcode/leetcode/p139_word-break/p139_2"
	"testing"
)

//func Test_1(t *testing.T) {
//	testWordBreak(t, p139_1.WordBreak)
//}

func Test_2(t *testing.T) {
	testWordBreak(t, p139_2.WordBreak)
}

func testWordBreak(t *testing.T, f func(string, []string) bool) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"sample-0", args{"leetcode", []string{"leet", "code"}}, true},
		{"sample-1", args{"applepenapple", []string{"apple", "pen"}}, true},
		{"sample-2", args{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}}, false},
		/*
			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
			["a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"]
		*/
		{"sample-3", args{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.s, tt.args.wordDict); got != tt.want {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
