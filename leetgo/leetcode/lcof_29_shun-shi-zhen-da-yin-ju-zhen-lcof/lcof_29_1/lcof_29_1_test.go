package lcof_29_1

import (
	"reflect"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"sample-0", args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{"sample-1", args{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
		{"case-0", args{[][]int{{1, 2, 3, 4}}}, []int{1, 2, 3, 4}},
		{"case-1", args{[][]int{{1}, {2}, {3}, {4}}}, []int{1, 2, 3, 4}},
		{"case-2", args{[][]int{{1, 2, 3}, {4, 5, 6}}}, []int{1, 2, 3, 6, 5, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralOrder(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
