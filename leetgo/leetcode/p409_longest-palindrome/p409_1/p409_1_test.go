package p409_1

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{"abccccdd"}, 7},
		{"case-0", args{"abcdef"}, 1},
		{"case-1", args{"abcdee"}, 3},
		{"case-2", args{"aabb"}, 4},
		{"case-3", args{"a"}, 1},
		{"case-4", args{""}, 0},
		{"case-5", args{"ab"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
