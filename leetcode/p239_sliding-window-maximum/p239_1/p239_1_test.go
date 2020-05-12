package p239_1

import (
	"reflect"
	"testing"
)

func Test_maxSlidingWindow(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"sample-0", args{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3}, []int{3, 3, 5, 5, 6, 7}},
		//[-7,-8,7,5,7,1,6,0]
		//4
		{"sample-1", args{[]int{-7, -8, 7, 5, 7, 1, 6, 0}, 4}, []int{7, 7, 7, 7, 7}},
		{"case-0", args{[]int{1, 3, -1, 0, 2, 1, 3, 2}, 3}, []int{3, 3, 2, 2, 3, 3}},
		{"case-1", args{[]int{1, 3, -1, 0, 4, 1, 3, 2}, 8}, []int{4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindow(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
