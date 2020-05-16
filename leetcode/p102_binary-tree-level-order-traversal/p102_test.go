package p102_binary_tree_level_order_traversal

import (
	"github.com/TeslaCN/goleetcode/leetcode/p102_binary-tree-level-order-traversal/p102_1"
	"github.com/TeslaCN/goleetcode/leetcode/p102_binary-tree-level-order-traversal/p102_2"
	"github.com/TeslaCN/goleetcode/util"
	"reflect"
	"testing"
)

type TreeNode = util.TreeNode

func Test_1(t *testing.T) {
	testLevelOrder(t, p102_1.LevelOrder)
}

func Test_2(t *testing.T) {
	testLevelOrder(t, p102_2.LevelOrder)
}

func testLevelOrder(t *testing.T, f func(*TreeNode) [][]int) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"sample-0", args{util.MakeTreeFromString("[3,9,20,null,null,15,7]")}, [][]int{{3}, {9, 20}, {15, 7}}},
		{"case-0", args{util.MakeTreeFromString("[3,9,20,-1,-2,15,7]")}, [][]int{{3}, {9, 20}, {-1, -2, 15, 7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
