package p31_1

import (
	"reflect"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name   string
		args   args
		expect []int
	}{
		{"sample-0", args{[]int{1, 2, 3}}, []int{1, 3, 2}},
		{"sample-1", args{[]int{3, 2, 1}}, []int{1, 2, 3}},
		{"sample-2", args{[]int{1, 1, 5}}, []int{1, 5, 1}},
		{"sample-3", args{[]int{1, 3, 2}}, []int{2, 1, 3}},
		{"sample-4", args{[]int{2, 3, 1, 3, 3}}, []int{2, 3, 3, 1, 3}},
		{"case-0", args{[]int{1, 3, 2, 1}}, []int{2, 1, 1, 3}},
		{"case-1", args{[]int{1, 2, 3, 4, 5, 4, 3, 2}}, []int{1, 2, 3, 5, 2, 3, 4, 4}},
		{"case-2", args{[]int{2, 3, 1}}, []int{3, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.expect) {
				t.Errorf("Expect: %v, Got: %v", tt.expect, tt.args.nums)
			}
		})
	}
}
