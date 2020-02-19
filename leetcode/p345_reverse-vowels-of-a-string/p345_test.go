package p345_reverse_vowels_of_a_string

import (
	"github.com/TeslaCN/goleetcode/leetcode/p345_reverse-vowels-of-a-string/p345_1"
	"testing"
)

type args struct {
	s string
}

var tests = []struct {
	name string
	args args
	want string
}{
	{"sample-0", args{"hello"}, "holle"},
	{"sample-1", args{"leetcode"}, "leotcede"},
	{"case-0", args{"aA"}, "Aa"},
}

func Test_reverseVowels(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p345_1.ReverseVowels(tt.args.s); got != tt.want {
				t.Errorf("reverseVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
