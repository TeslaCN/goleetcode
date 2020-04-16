package p56_1

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"sample-0", args{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{"sample-1", args{[][]int{{1, 4}, {4, 5}}}, [][]int{{1, 5}}},
		{"case-0", args{[][]int{{5, 9}, {2, 6}, {0, 1}, {2, 4}, {10, 18}}}, [][]int{{0, 1}, {2, 9}, {10, 18}}},
		{"case-1", args{[][]int{}}, [][]int{}},
		{"case-2", args{[][]int{{0, 1}}}, [][]int{{0, 1}}},
		{"case-3", args{[][]int{{0, 1}, {2, 3}}}, [][]int{{0, 1}, {2, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
