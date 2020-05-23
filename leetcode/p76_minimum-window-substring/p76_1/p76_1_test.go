package p76_1

import "testing"

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"sample-0", args{"ADOBECODEBANC", "ABC"}, "BANC"},
		{"case-0", args{"CNABEDOCEBODA", "ABC"}, "CNAB"},
		{"case-1", args{"DNABEDOBEBODA", "ABC"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
