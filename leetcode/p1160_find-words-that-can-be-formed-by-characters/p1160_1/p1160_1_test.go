package p1160_1

import "testing"

func Test_countCharacters(t *testing.T) {
	type args struct {
		words []string
		chars string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{[]string{"cat", "bt", "hat", "tree"}, "atach"}, 6},
		{"sample-1", args{[]string{"hello", "world", "leetcode"}, "welldonehoneyr"}, 10},
		{"case-0", args{[]string{""}, "welldonehoneyr"}, 0},
		{"case-1", args{[]string{"a", "e"}, "welldonehoneyr"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCharacters(tt.args.words, tt.args.chars); got != tt.want {
				t.Errorf("countCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
