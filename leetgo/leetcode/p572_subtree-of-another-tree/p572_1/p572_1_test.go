package p572_1

import (
	"github.com/TeslaCN/goleetcode/util"
	"testing"
)

func Test_isSubtree(t *testing.T) {
	type args struct {
		s *TreeNode
		t *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"sample-0", args{util.MakeTreeFromString("[3,4,5,1,2]"), util.MakeTreeFromString("[4,1,2]")}, true},
		{"sample-1", args{util.MakeTreeFromString("[3,4,5,1,2,null,null,null,null,0]"), util.MakeTreeFromString("[4,1,2]")}, false},
		{"sample-2", args{util.MakeTreeFromString("[3,4,5,1,2,null,null,0]"), util.MakeTreeFromString("[4,1,2]")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubtree(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}
