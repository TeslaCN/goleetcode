package p347_1

import (
	"reflect"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"sample-2", args{[]int{3, 0, 1, 0}, 1}, []int{0}},
		{"case-0", args{[]int{1, 1, 1, 2, 2, 4, 4, 4, 4, 4, 5, 5, 5, 5, 3}, 3}, []int{4, 5, 1}},
		{"sample-0", args{[]int{1, 1, 1, 2, 2, 3}, 2}, []int{1, 2}},
		{"sample-1", args{[]int{1}, 1}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
