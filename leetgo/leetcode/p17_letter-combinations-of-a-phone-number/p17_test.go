package p17_letter_combinations_of_a_phone_number

import (
	"github.com/TeslaCN/goleetcode/leetcode/p17_letter-combinations-of-a-phone-number/p17_1"
	"reflect"
	"testing"
)

type args struct {
	digits string
}

var tests = []struct {
	name string
	args args
	want []string
}{
	{"sample-0", args{"23"}, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
}

func Test_letterCombinations(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p17_1.LetterCombinations(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}
