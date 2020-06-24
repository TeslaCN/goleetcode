package p438_find_all_anagrams_in_a_string

import (
	"github.com/TeslaCN/goleetcode/leetcode/p438_find-all-anagrams-in-a-string/p438_1"
	"github.com/TeslaCN/goleetcode/leetcode/p438_find-all-anagrams-in-a-string/p438_2"
	"reflect"
	"testing"
)

func Test_1(t *testing.T) {
	testFindAnagrams(t, p438_1.FindAnagrams)
}

func Test_2(t *testing.T) {
	testFindAnagrams(t, p438_2.FindAnagrams)
}

func testFindAnagrams(t *testing.T, f func(string, string) []int) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"sample-0", args{"cbaebabacd", "abc"}, []int{0, 6}},
		{"sample-1", args{"abab", "ab"}, []int{0, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.s, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
