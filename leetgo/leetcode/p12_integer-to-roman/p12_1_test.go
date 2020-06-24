package p12_integer_to_roman

import (
	"github.com/TeslaCN/goleetcode/leetcode/p12_integer-to-roman/p12_1"
	"testing"
)

type args struct {
	num int
}

var tests = []struct {
	name string
	args args
	want string
}{
	{"sample-0", args{3}, "III"},
	{"sample-1", args{4}, "IV"},
	{"sample-2", args{9}, "IX"},
	{"sample-3", args{58}, "LVIII"},
	{"sample-4", args{1994}, "MCMXCIV"},
	{"case-0", args{3999}, "MMMCMXCIX"},
}

func Test_intToRoman(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p12_1.IntToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
