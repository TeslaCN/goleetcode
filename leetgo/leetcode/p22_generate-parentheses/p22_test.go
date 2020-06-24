package p22_generate_parentheses

import (
	"github.com/TeslaCN/goleetcode/leetcode/p22_generate-parentheses/p22_1"
	"github.com/TeslaCN/goleetcode/leetcode/p22_generate-parentheses/p22_2"
	"github.com/TeslaCN/goleetcode/leetcode/p22_generate-parentheses/p22_3"
	"reflect"
	"sort"
	"testing"
)

func Test_1(t *testing.T) {
	testGenerateParenthesis(t, p22_1.GenerateParenthesis)
}
func Test_2(t *testing.T) {
	testGenerateParenthesis(t, p22_2.GenerateParenthesis)
}

func Test_3(t *testing.T) {
	testGenerateParenthesis(t, p22_3.GenerateParenthesis)
}
func testGenerateParenthesis(t *testing.T, f func(int) []string) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"sample-0", args{3}, []string{
			"((()))",
			"(()())",
			"(())()",
			"()(())",
			"()()()",
		}},
		{"sample-1", args{4}, []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"}},
		{"case-0", args{2}, []string{
			"(())",
			"()()",
		}},
		{"case-1", args{1}, []string{
			"()",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := f(tt.args.n)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}
