package p990_satisfiability_of_equality_equations

import (
	"github.com/TeslaCN/goleetcode/leetcode/p990_satisfiability-of-equality-equations/p990_1"
	"github.com/TeslaCN/goleetcode/leetcode/p990_satisfiability-of-equality-equations/p990_2"
	"testing"
)

func Test_1(t *testing.T) {
	testEquationsPossible(t, p990_1.EquationsPossible)
}

func Test_2(t *testing.T) {
	testEquationsPossible(t, p990_2.EquationsPossible)
}

func testEquationsPossible(t *testing.T, f func([]string) bool) {
	type args struct {
		equations []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"sample-0", args{[]string{"a==b", "b!=a"}}, false},
		{"sample-1", args{[]string{"a==b", "b==a"}}, true},
		{"sample-2", args{[]string{"a==b", "b==c", "a==c"}}, true},
		{"sample-3", args{[]string{"a==b", "b!=c", "c==a"}}, false},
		{"sample-4", args{[]string{"c==c", "b==d", "x!=z"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.equations); got != tt.want {
				t.Errorf("equationsPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
