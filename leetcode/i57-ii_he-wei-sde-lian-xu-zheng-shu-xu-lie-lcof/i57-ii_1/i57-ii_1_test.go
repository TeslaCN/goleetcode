package i57_ii_1

import (
	"reflect"
	"testing"
)

func Test_findContinuousSequence(t *testing.T) {
	type args struct {
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"sample-0", args{9}, [][]int{{2, 3, 4}, {4, 5}}},
		{"sample-1", args{15}, [][]int{{1, 2, 3, 4, 5}, {4, 5, 6}, {7, 8}}},
		{"case-0", args{1}, [][]int{}},
		{"case-1", args{2}, [][]int{}},
		{"case-2", args{3}, [][]int{{1, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findContinuousSequence(tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findContinuousSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
