package p114_flatten_binary_tree_to_linked_list

import (
	"github.com/TeslaCN/goleetcode/leetcode/p114_flatten-binary-tree-to-linked-list/p114_1"
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
	{"sample-0", args{util.MakeTreeFromString("[1,2,5,3,4,null,6]")}, util.MakeTreeFromString("[1,null,2,null,3,null,4,null,5,null,6]")},
	{"case-0", args{util.MakeTreeFromString("[1,2,3]")}, util.MakeTreeFromString("[1,null,2,null,3]")},
}

func Test_flatten(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p114_1.Flatten(tt.args.root)
			if !reflect.DeepEqual(tt.args.root, tt.want) {
				t.Errorf("flatten() = %v, want %v", tt.args.root, tt.want)
			}
		})
	}
}
