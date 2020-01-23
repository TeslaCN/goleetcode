package p538_convert_bst_to_greater_tree

import (
	"github.com/TeslaCN/goleetcode/leetcode/p538_convert-bst-to-greater-tree/p538_1"
	"github.com/TeslaCN/goleetcode/util"
	"reflect"
	"testing"
)

type TreeNode = util.TreeNode

type args struct {
	root *TreeNode
}

var tests = []struct {
	name string
	args args
	want *TreeNode
}{
	{
		/*
			     5
			   /   \
			  2     13

				 18
				/   \
			  20     13
		*/
		name: "sample-0",
		args: args{util.CreateTreeFromSequence(util.ConvertToIntPointer([]int{5, 2, 13}))},
		want: util.CreateTreeFromSequence(util.ConvertToIntPointer([]int{18, 20, 13})),
	},
	{
		/*
					5
			      /   \
			     3      8
			    / \    /  \
			   2   4  6    9
			  /        \    \
			 1          7    10

					  45
			       /      \
			     52        27
			    /  \      /   \
			   54   49   40    19
			  /            \     \
			 55             34    10
		*/
		name: "case-0",
		args: args{util.CreateTreeFromSequenceInt([]int{5, 3, 8, 2, 4, 6, 9, 1, 0, 0, 0, 0, 7, 0, 10})},
		want: util.CreateTreeFromSequenceInt([]int{45, 52, 27, 54, 49, 40, 19, 55, 0, 0, 0, 0, 34, 0, 10}),
	},
	{
		/*
		        5
		     2     13
		   1  3   8  15

		          41
		      46       28
		   47   44   36   15
		*/
		name: "case-1",
		args: args{util.CreateTreeFromSequenceInt([]int{5, 2, 13, 1, 3, 8, 15})},
		want: util.CreateTreeFromSequenceInt([]int{41, 46, 28, 47, 44, 36, 15}),
	},
}

func TestConvertBST(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p538_1.ConvertBST(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
