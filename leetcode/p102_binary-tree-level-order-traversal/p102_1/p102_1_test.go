package p102_1

import (
	"github.com/TeslaCN/goleetcode/util"
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
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
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
