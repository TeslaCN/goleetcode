package p660_valid_palindrome_ii

import (
	"github.com/TeslaCN/goleetcode/leetcode/p660_valid-palindrome-ii/p660_1"
	"testing"
)

type args struct {
	s string
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{"sample-3", args{"aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"}, true},
	{"sample-0", args{"aba"}, true},
	{"sample-1", args{"abca"}, true},
	{"sample-2", args{"abc"}, false},
	{"case-0", args{"abcdecba"}, true},
	{"case-1", args{"abcecba"}, true},
	{"case-2", args{"abcdefgdcba"}, false},
	{"case-3", args{"ab"}, true},
	{"case-4", args{"cbab"}, true},
	{"case-5", args{"bcab"}, true},
	{"case-6", args{"acbaba"}, true},
	{"case-7", args{"abcaba"}, true},
	{"case-8", args{"abcdaacba"}, true},
	{"case-9", args{"lcuucul"}, true},
}

func Test_validPalindrome(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p660_1.ValidPalindrome(tt.args.s); got != tt.want {
				t.Errorf("validPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
