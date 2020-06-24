package p46_permutations

import (
	"github.com/TeslaCN/goleetcode/leetcode/p46_permutations/p46_1"
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
	{"sample-0", args{[]int{1, 2, 3}}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
	{"sample-1", args{[]int{1, 2, 3, 4}}, [][]int{{1, 2, 3, 4}, {1, 2, 4, 3}, {1, 3, 2, 4}, {1, 3, 4, 2}, {1, 4, 2, 3}, {1, 4, 3, 2}, {2, 1, 3, 4}, {2, 1, 4, 3}, {2, 3, 1, 4}, {2, 3, 4, 1}, {2, 4, 1, 3}, {2, 4, 3, 1}, {3, 1, 2, 4}, {3, 1, 4, 2}, {3, 2, 1, 4}, {3, 2, 4, 1}, {3, 4, 1, 2}, {3, 4, 2, 1}, {4, 1, 2, 3}, {4, 1, 3, 2}, {4, 2, 1, 3}, {4, 2, 3, 1}, {4, 3, 1, 2}, {4, 3, 2, 1}}},
	{"case-0", args{[]int{1, 2}}, [][]int{{1, 2}, {2, 1}}},
}

func Test_permute(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p46_1.Permute(tt.args.nums); !util.EqualsIgnoreOrder(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}
