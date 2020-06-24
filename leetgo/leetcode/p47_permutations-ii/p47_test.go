package p47_permutations_ii

import (
	"github.com/TeslaCN/goleetcode/leetcode/p47_permutations-ii/p47_1"
	"github.com/TeslaCN/goleetcode/util"
	"testing"
)

type args struct {
	nums []int
}

var tests = []struct {
	name string
	args args
	want [][]int
}{
	{"sample-0", args{[]int{1, 1, 2}}, [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
	{"sample-1", args{[]int{2, 2, 1, 1}}, [][]int{{1, 1, 2, 2}, {1, 2, 1, 2}, {1, 2, 2, 1}, {2, 1, 1, 2}, {2, 1, 2, 1}, {2, 2, 1, 1}}},
	{"sample-2", args{[]int{3, 3, 0, 3}}, [][]int{{0, 3, 3, 3}, {3, 0, 3, 3}, {3, 3, 0, 3}, {3, 3, 3, 0}}},
	{"case-0", args{[]int{1, 2, 3}}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
	{"case-1", args{[]int{1, 2, 3, 4}}, [][]int{{1, 2, 3, 4}, {1, 2, 4, 3}, {1, 3, 2, 4}, {1, 3, 4, 2}, {1, 4, 2, 3}, {1, 4, 3, 2}, {2, 1, 3, 4}, {2, 1, 4, 3}, {2, 3, 1, 4}, {2, 3, 4, 1}, {2, 4, 1, 3}, {2, 4, 3, 1}, {3, 1, 2, 4}, {3, 1, 4, 2}, {3, 2, 1, 4}, {3, 2, 4, 1}, {3, 4, 1, 2}, {3, 4, 2, 1}, {4, 1, 2, 3}, {4, 1, 3, 2}, {4, 2, 1, 3}, {4, 2, 3, 1}, {4, 3, 1, 2}, {4, 3, 2, 1}}},
	{"case-2", args{[]int{1, 2}}, [][]int{{1, 2}, {2, 1}}},
	{"case-3", args{[]int{1, 1, 1}}, [][]int{{1, 1, 1}}},
}

func Test_permuteUnique(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p47_1.PermuteUnique(tt.args.nums); !util.EqualsIgnoreOrder(got, tt.want) {
				t.Errorf("permuteUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
