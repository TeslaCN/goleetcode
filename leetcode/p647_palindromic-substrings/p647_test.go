package p647_palindromic_substrings

import (
	"github.com/TeslaCN/goleetcode/leetcode/p647_palindromic-substrings/p647_1"
	"github.com/TeslaCN/goleetcode/leetcode/p647_palindromic-substrings/p647_2"
	"testing"
)

func Test_1(t *testing.T) {
	testCountSubstrings(t, p647_1.CountSubstrings)
}

func Test_2(t *testing.T) {
	testCountSubstrings(t, p647_2.CountSubstrings)
}

func testCountSubstrings(t *testing.T, f func(string) int) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{"abc"}, 3},
		{"sample-1", args{"aaa"}, 6},
		{"case-0", args{"aaba"}, 6},
		{"case-1", args{"abba"}, 6},
		{"case-2", args{"aba"}, 4},
		{"case-3", args{"abacba"}, 7},
		{"case-4", args{"abaaba"}, 11},
		{"case-5", args{"abacaba"}, 7 + 3 + 1 + 1},
		{"case-6", args{"bbaaba"}, 6 + 2 + 1 + 1},
		{"case-7", args{"cabaaba"}, 7 + 1 + 2 + 1 + 1},
		{"case-8", args{"cabacaba"}, 8 + 3 + 2 + 1},
		{"case-9", args{"abcba"}, 5 + 1 + 1},
		{"case-large", args{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}, 500500},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.s); got != tt.want {
				t.Errorf("countSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
