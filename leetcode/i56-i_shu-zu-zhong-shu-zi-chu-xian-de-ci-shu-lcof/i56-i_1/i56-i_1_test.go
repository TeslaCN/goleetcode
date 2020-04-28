package i56_i_1

import (
	"reflect"
	"sort"
	"testing"
)

func Test_singleNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"sample-0", args{[]int{4, 1, 4, 6}}, []int{1, 6}},
		{"sample-1", args{[]int{1, 2, 10, 4, 1, 4, 3, 3}}, []int{2, 10}},
		{"case-0", args{[]int{1, 5, 10, 4, 1, 4, 3, 3}}, []int{5, 10}},
		{"case-1", args{[]int{1, 6, 4, 1, 4, 10, 3, 3}}, []int{6, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := singleNumbers(tt.args.nums)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("singleNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
