package p617_merge_two_binary_trees

import (
	"github.com/TeslaCN/goleetcode/leetcode/p617_merge-two-binary-trees/p617_1"
	"github.com/TeslaCN/goleetcode/leetcode/p617_merge-two-binary-trees/p617_2"
	"github.com/TeslaCN/goleetcode/util"
	"reflect"
	"testing"
)

type TreeNode = util.TreeNode

type args struct {
	t1 *TreeNode
	t2 *TreeNode
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
		{"sample-0", args{
			t1: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 5},
				},
				Right: &TreeNode{Val: 2},
			},
			t2: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   1,
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 7},
				},
			},
		}, &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 5},
				Right: &TreeNode{Val: 4},
			},
			Right: &TreeNode{
				Val:   5,
				Right: &TreeNode{Val: 7},
			},
		}},
		/*
				    1          1
				   /		    \
				  2              2
				 /                \
				3                  3

			    2
			   / \
			  2   2
			 /     \
			3       3
		*/
		{"sample-1", args{
			t1: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
			},
			t2: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
			},
		}, &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 3},
			},
			Right: &TreeNode{
				Val:   2,
				Right: &TreeNode{Val: 3},
			},
		}},
		{"case-0", args{&TreeNode{Val: 1}, &TreeNode{Val: 9}}, &TreeNode{Val: 10}},
	}...)
}

func Test_1(t *testing.T) {
	testMergeTrees(t, p617_1.MergeTrees)
}

func Test_2(t *testing.T) {
	testMergeTrees(t, p617_2.MergeTrees)
}

func testMergeTrees(t *testing.T, f func(t1, t2 *TreeNode) *TreeNode) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.t1, tt.args.t2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeTrees() = %v, want %v",
					util.FormatIntPointerSlice(util.SequenceTraversalValues(got)),
					util.FormatIntPointerSlice(util.SequenceTraversalValues(tt.want)),
				)
			}
		})
	}
}
