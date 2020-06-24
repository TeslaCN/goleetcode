package p105_1

import (
	"github.com/TeslaCN/goleetcode/util"
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"sample-0", args{[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}}, util.MakeTreeFromString("[3,9,20,null,null,15,7]")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.preorder, tt.args.inorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
