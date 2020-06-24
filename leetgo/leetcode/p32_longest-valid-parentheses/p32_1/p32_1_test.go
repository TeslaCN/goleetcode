package p32_1

import "testing"

func Test_longestValidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-4", args{"))))())()()(()"}, 4},
		{"sample-0", args{"(()"}, 2},
		{"sample-1", args{")()())"}, 4},
		{"sample-2", args{")("}, 0},
		{"sample-3", args{"))))((()(("}, 2},
		{"case-0", args{")())())"}, 2},
		{"case-1", args{"(()(()))"}, 8},
		{"case-2", args{")((()))"}, 6},
		{"case-3", args{")()("}, 2},
		{"case-4", args{"())(()"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParentheses(tt.args.s); got != tt.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
