package p98_1

import (
	"github.com/TeslaCN/goleetcode/util"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"sample-0", args{util.MakeTreeFromString("[2,1,3]")}, true},
		{"sample-1", args{util.MakeTreeFromString("[5,1,4,null,null,3,6]")}, false},
		{"case-0", args{util.MakeTreeFromString("[]")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
