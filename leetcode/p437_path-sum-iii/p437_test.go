package p437_path_sum_iii

import (
	"github.com/TeslaCN/goleetcode/leetcode/p437_path-sum-iii/p437_1"
	"github.com/TeslaCN/goleetcode/util"
	"testing"
)

type TreeNode = util.TreeNode

type args struct {
	root *TreeNode
	sum  int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{
		name: "sample-0",
		/*
			      10
			     /  \
			    5   -3
			   / \    \
			  3   2   11
			 / \   \
			3  -2   1
		*/
		args: args{
			root: &p437_1.TreeNode{
				Val: 10,
				Left: &p437_1.TreeNode{
					Val: 5,
					Left: &p437_1.TreeNode{
						Val: 3,
						Left: &p437_1.TreeNode{
							Val:  3,
							Left: &p437_1.TreeNode{Val: 3},
						},
						Right: &p437_1.TreeNode{Val: -2},
					},
					Right: &p437_1.TreeNode{
						Val:   2,
						Right: &p437_1.TreeNode{Val: 1},
					},
				},
				Right: &p437_1.TreeNode{
					Val:   -3,
					Right: &p437_1.TreeNode{Val: 11},
				},
			},
			sum: 8,
		},
		want: 3,
	},
	{
		name: "case-0",
		/*
				      10
				     /  \
				    5   -3
				   / \    \
				  3   2   11
				 / \   \    \
				3  -2   1   -11
			                /
			               1
		*/
		args: args{
			root: &p437_1.TreeNode{
				Val: 10,
				Left: &p437_1.TreeNode{
					Val: 5,
					Left: &p437_1.TreeNode{
						Val: 3,
						Left: &p437_1.TreeNode{
							Val:  3,
							Left: &p437_1.TreeNode{Val: 3},
						},
						Right: &p437_1.TreeNode{Val: -2},
					},
					Right: &p437_1.TreeNode{
						Val:   2,
						Right: &p437_1.TreeNode{Val: 1},
					},
				},
				Right: &p437_1.TreeNode{
					Val: -3,
					Right: &p437_1.TreeNode{
						Val: 11,
						Right: &p437_1.TreeNode{
							Val:  -11,
							Left: &p437_1.TreeNode{Val: 1},
						},
					},
				},
			},
			sum: 8,
		},
		want: 4,
	},
	{
		name: "case-1",
		args: args{
			root: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val:   8,
					Left:  &TreeNode{Val: 8},
					Right: &TreeNode{Val: 8},
				},
				Right: &TreeNode{
					Val:   8,
					Left:  &TreeNode{Val: 8},
					Right: &TreeNode{Val: 8},
				},
			},
			sum: 8,
		},
		want: 7,
	},
}

func TestPathSum(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p437_1.PathSum(tt.args.root, tt.args.sum); got != tt.want {
				t.Errorf("PathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
