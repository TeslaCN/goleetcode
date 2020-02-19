package p88_merge_sorted_array

import (
	"github.com/TeslaCN/goleetcode/leetcode/p88_merge-sorted-array/p88_1"
	"reflect"
	"testing"
)

func Test_1(t *testing.T) {
	testMerge(t, p88_1.Merge)
}

func testMerge(t *testing.T, f func([]int, int, []int, int)) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "sample-1",
			args: args{
				nums1: []int{2, 0},
				m:     1,
				nums2: []int{1},
				n:     1,
			},
			want: []int{1, 2},
		},
		{
			name: "sample-0",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
			want: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name: "case-0",
			args: args{
				nums1: []int{1, 2, 3},
				m:     3,
				nums2: []int{},
				n:     0,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "case-1",
			args: args{
				nums1: []int{0, 0, 0},
				m:     0,
				nums2: []int{1, 2, 3},
				n:     3,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "case-2",
			args: args{
				nums1: []int{1, 0, 0, 0},
				m:     1,
				nums2: []int{1, 2, 3},
				n:     3,
			},
			want: []int{1, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			if !reflect.DeepEqual(tt.args.nums1, tt.want) {
				t.Errorf("flatten() = %v, want %v", tt.args.nums1, tt.want)
			}
		})
	}
}
