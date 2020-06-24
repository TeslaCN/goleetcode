package p78_subsets

import (
	"github.com/TeslaCN/goleetcode/leetcode/p78_subsets/p78_1"
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
	{"sample-7", args{[]int{5, 1, 2, 3, 4}}, [][]int{{}, {5}, {1}, {1, 5}, {2}, {2, 5}, {1, 2}, {1, 2, 5}, {3}, {3, 5}, {1, 3}, {1, 3, 5}, {2, 3}, {2, 3, 5}, {1, 2, 3}, {1, 2, 3, 5}, {4}, {4, 5}, {1, 4}, {1, 4, 5}, {2, 4}, {2, 4, 5}, {1, 2, 4}, {1, 2, 4, 5}, {3, 4}, {3, 4, 5}, {1, 3, 4}, {1, 3, 4, 5}, {2, 3, 4}, {2, 3, 4, 5}, {1, 2, 3, 4}, {1, 2, 3, 4, 5}}},
	{"sample-0", args{[]int{1, 2, 3}}, [][]int{{3}, {1}, {2}, {1, 2, 3}, {1, 3}, {2, 3}, {1, 2}, {}}},
	{"case-0", args{[]int{1, 2}}, [][]int{{1}, {2}, {1, 2}, {}}},
}

func Test_subsets(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p78_1.Subsets(tt.args.nums); !util.EqualsIgnoreOrder(got, tt.want) {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOutOfBoundary(t *testing.T) {
	t.Log([]int{1, 2, 3}[1:])
	t.Log([]int{1, 2, 3}[3:])
}
