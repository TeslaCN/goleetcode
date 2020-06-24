package i01_06

import "testing"

func Test_compressString(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"sample-0", args{"aabcccccaaa"}, "a2b1c5a3"},
		{"sample-1", args{"abbccd"}, "abbccd"},
		{"sample-2", args{"a"}, "a"},
		{"sample-3", args{"bb"}, "bb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compressString(tt.args.S); got != tt.want {
				t.Errorf("compressString() = %v, want %v", got, tt.want)
			}
		})
	}
}
