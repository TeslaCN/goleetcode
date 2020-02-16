package p226_invert_binary_tree

import (
	"github.com/TeslaCN/goleetcode/leetcode/p226_invert-binary-tree/p226_1"
	"github.com/TeslaCN/goleetcode/util"
	"math"
	"reflect"
	"testing"
)

type TreeNode = util.TreeNode

type args struct {
	root *TreeNode
}

var tests []struct {
	name string
	args args
	want *TreeNode
}

func init() {
	tests = append(tests, []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"sample-0",
			args{util.CreateTreeFromSequence(util.ConvertToIntPointer([]int{4, 2, 7, 1, 3, 6, 9}))},
			util.CreateTreeFromSequence(util.ConvertToIntPointer([]int{4, 7, 2, 9, 6, 3, 1})),
		},
		{
			"case-0",
			args{util.CreateTreeFromSequence(util.ConvertToIntPointer([]int{4, 2, 7, 1, math.MinInt32, math.MinInt32, 9}))},
			util.CreateTreeFromSequence(util.ConvertToIntPointer([]int{4, 7, 2, 9, math.MinInt32, math.MinInt32, 1})),
		},
	}...)
}

func TestInvertTree(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p226_1.InvertTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InvertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
