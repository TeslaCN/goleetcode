package p543_diameter_of_binary_tree

import (
	"github.com/TeslaCN/goleetcode/leetcode/p543_diameter-of-binary-tree/p543_1"
	"github.com/TeslaCN/goleetcode/util"
	"testing"
)

type TreeNode = util.TreeNode

type args struct {
	root *TreeNode
}

var tests = []struct {
	name string
	args args
	want int
}{
	{
		name: "sample-0",
		want: 3,
		args: args{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{Val: 3},
			},
		},
	},
	{
		name: "sample-103",
		want: 8,
		args: args{
			root: &TreeNode{
				Val:  4,
				Left: &TreeNode{Val: -7},
				Right: &TreeNode{
					Val: -3,
					Left: &TreeNode{
						Val: -9,
						Left: &TreeNode{
							Val: 9,
							Left: &TreeNode{
								Val: 6,
								Left: &TreeNode{
									Val:  0,
									Left: nil,
									Right: &TreeNode{
										Val: -1,
									},
								},
								Right: &TreeNode{
									Val: 6,
									Left: &TreeNode{
										Val: -4,
									},
									Right: nil,
								},
							},
							Right: nil,
						},
						Right: &TreeNode{
							Val: -7,
							Left: &TreeNode{
								Val: -6,
								Left: &TreeNode{
									Val: 5,
								},
								Right: nil,
							},
							Right: &TreeNode{
								Val: -6,
								Left: &TreeNode{
									Val: 9,
									Left: &TreeNode{
										Val: -2,
									},
								},
								Right: nil,
							},
						},
					},
					Right: &TreeNode{
						Val:   -3,
						Left:  &TreeNode{Val: -4},
						Right: nil,
					},
				},
			},
		},
	},
	{
		name: "case-0",
		want: 2,
		args: args{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
			},
		},
	},
	{
		name: "case-1",
		want: 3,
		args: args{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
					Right: &TreeNode{
						Val:  4,
						Left: &TreeNode{Val: 5},
					},
				},
			},
		},
	},
	{
		name: "case-2",
		want: 2,
		args: args{
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
		},
	},
}

func Test_diameterOfBinaryTree(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p543_1.DiameterOfBinaryTree(tt.args.root); got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
