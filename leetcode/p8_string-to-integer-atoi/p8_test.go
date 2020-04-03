package p8_string_to_integer_atoi

import (
	"github.com/TeslaCN/goleetcode/leetcode/p8_string-to-integer-atoi/p8_1"
	"math"
	"testing"
)

func Test_1(t *testing.T) {
	testMyAtoi(t, p8_1.MyAtoi)
}

func testMyAtoi(t *testing.T, f func(string) int) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{"42"}, 42},
		{"sample-1", args{"   -42"}, -42},
		{"sample-2", args{"4193 with words"}, 4193},
		{"sample-3", args{"words and 987"}, 0},
		{"sample-4", args{"-91283472332"}, math.MinInt32},
		{"case-0", args{"991283472332"}, math.MaxInt32},
		{"case-1", args{"2147483648"}, math.MaxInt32},
		{"case-2", args{"2147483646"}, 2147483646},
		{"case-3", args{"-2147483648"}, math.MinInt32},
		{"case-4", args{"--2147483648"}, 0},
		{"case-5", args{" -2147483648"}, math.MinInt32},
		{"case-6", args{""}, 0},
		{"case-7", args{"   g1"}, 0},
		{"case-8", args{"   -1"}, -1},
		{"case-9", args{" 1234h5"}, 1234},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
