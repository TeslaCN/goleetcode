package p1071_1

import "testing"

func Test_gcdOfStrings(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"sample-0", args{"ABCABC", "ABC"}, "ABC"},
		{"sample-1", args{"ABABAB", "ABAB"}, "AB"},
		{"sample-2", args{"LEET", "CODE"}, ""},
		{"case-0", args{"AGCCT", "AGCCTAGCCT"}, "AGCCT"},
		{"case-1", args{"AAAAAA", "A"}, "A"},
		{"case-2", args{"AAAAAA", "AAA"}, "AAA"},
		{"case-3", args{"ABABAB", "ABA"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcdOfStrings(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("gcdOfStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
