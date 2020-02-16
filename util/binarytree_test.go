package util

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

/*
	     3
	    / \
	   4   5
	  / \   \
	 5   4   7

	0  1  2  3  4   5   6
	3, 4, 5, 5, 4, nil, 7
	0  1  1  2  2   2   2
	2n+1
*/
var root = &TreeNode{
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
}

func TestSequenceTraversal(t *testing.T) {
	nodes := SequenceTraversal(root)
	t.Logf("%v", nodes)
	values := SequenceTraversalValues(root)
	t.Log(FormatIntPointerSlice(values))
}

func TestCreateTreeFromSequence(t *testing.T) {
	//nums := []int{3, 4, 5, 5, 4, -1, 7}
	nums := []int{1, -1, 2, -1, 3}
	size := len(nums)
	inputs := make([]*int, size)
	for i := 0; i < size; i++ {
		if nums[i] != -1 {
			inputs[i] = &nums[i]
		}
	}
	tree := CreateTreeFromSequence(inputs)
	t.Logf("DeepEqual: %v", reflect.DeepEqual(root, tree))
}

func TestPreorderTraversal(t *testing.T) {
	PreorderTraversal(root, nil)
	PreorderTraversal(root, func(node *TreeNode) {
		fmt.Println(node)
	})
}

func TestInorderTraversal(t *testing.T) {
	InorderTraversal(root, nil)
	InorderTraversal(root, func(node *TreeNode) {
		fmt.Println(node)
	})
}

func TestPostorderTraversal(t *testing.T) {
	PostorderTraversal(root, nil)
	PostorderTraversal(root, func(node *TreeNode) {
		fmt.Println(node)
	})
}

func TestCreateTreeFromSequence1(t *testing.T) {
	type args struct {
		values []*int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "case-0",
			args: args{ConvertToIntPointer([]int{1, math.MinInt32, 2, math.MinInt32, 3})},
			want: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateTreeBySequenceTraversal(tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTreeFromSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeTreeFromString(t *testing.T) {
	MakeTreeFromString("[4,-7,-3,null,null,-9,-3,9,-7,-4,null,6,null,-6,-6,null,null,0,6,5,null,9,null,null,-1,-4,null,null,null,-2]")
}

func TestMakeTreeFromString1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "sample-0",
			args: args{"  [   1,  null,2, null,  3 ,null,4,null,5,  null ,   6  ] "},
			want: &TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val:  2,
					Left: nil,
					Right: &TreeNode{
						Val:  3,
						Left: nil,
						Right: &TreeNode{
							Val:  4,
							Left: nil,
							Right: &TreeNode{
								Val:   5,
								Left:  nil,
								Right: &TreeNode{Val: 6},
							},
						},
					},
				},
			},
		},
		{
			name: "sample-1",
			args: args{" [1, 2, 5, 3,4,  null ,6 ]    "},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: &TreeNode{Val: 6},
				},
			},
		},
		{
			name: "sample-2",
			args: args{"[1, 2, null, 3]"},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3},
					Right: nil,
				},
				Right: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeTreeFromString(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeTreeFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
