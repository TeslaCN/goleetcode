package p820_short_encoding_of_words

import (
	"github.com/TeslaCN/goleetcode/leetcode/p820_short-encoding-of-words/p820_1"
	"github.com/TeslaCN/goleetcode/leetcode/p820_short-encoding-of-words/p820_2"
	"github.com/TeslaCN/goleetcode/leetcode/p820_short-encoding-of-words/p820_3"
	"testing"
)

func Test_1(t *testing.T) {
	testMinimumLengthEncoding(t, p820_1.MinimumLengthEncoding)
}

func Test_2(t *testing.T) {
	testMinimumLengthEncoding(t, p820_2.MinimumLengthEncoding)
}

func Test_3(t *testing.T) {
	testMinimumLengthEncoding(t, p820_3.MinimumLengthEncoding)
}

func testMinimumLengthEncoding(t *testing.T, f func([]string) int) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{[]string{"time", "me", "bell"}}, 10},
		{"sample-1", args{[]string{"me", "time"}}, 5},
		{"sample-2", args{[]string{"atime", "time", "btime"}}, 12},
		{"case-0", args{[]string{"time"}}, 5},
		{"case-1", args{[]string{"time", "me"}}, 5},
		{"case-2", args{[]string{"time", "time"}}, 5},
		{"case-3", args{[]string{"time", "apple", "delta"}}, 17},
		{"case-4", args{[]string{"time", "apple", "apple", "delta"}}, 17},
		{"case-5", args{[]string{"time", "apple", "ple", "e"}}, 11},
		{"case-6", args{[]string{""}}, 1},
		{"case-7", args{[]string{"time", "time", "time"}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.words); got != tt.want {
				t.Errorf("minimumLengthEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}
