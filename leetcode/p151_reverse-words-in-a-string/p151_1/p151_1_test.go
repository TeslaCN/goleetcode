package p151_1

import "testing"

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"sample-2", args{"a good   example"}, "example good a"},
		{"sample-0", args{"the sky is blue"}, "blue is sky the"},
		{"sample-1", args{"  hello world!"}, "world! hello"},
		{"case-0", args{"   111    22 666666 333 4444  55555     "}, "55555 4444 333 666666 22 111"},
		{"case-1", args{"   11111    22 333 4444  55555     "}, "55555 4444 333 22 11111"},
		{"case-2", args{"   12321  "}, "12321"},
		{"case-3", args{""}, ""},
		{"case-4", args{"    a  "}, "a"},
		{"case-5", args{"a"}, "a"},
		{"case-6", args{"   a"}, "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
