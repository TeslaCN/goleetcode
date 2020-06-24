package p72_1

import "testing"

func Test_minDistance(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case-1", args{"a", "b"}, 1},
		{"case-0", args{"word1", "word2"}, 1},
		{"sample-0", args{"horse", "ros"}, 3},
		{"sample-1", args{"intention", "execution"}, 5},
		{"case-2", args{"a", "a"}, 0},
		{"case-3", args{"", ""}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistance(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("minDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
