package p84_1

import "testing"

func Test_largestRectangleArea(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{[]int{2, 1, 5, 6, 2, 3}}, 10},
		{"case-0", args{[]int{2, 2, 4, 6, 2, 3}}, 12},
		{"case-1", args{[]int{7, 8, 4, 1, 2, 3, 5, 9}}, 14},
		{"case-2", args{[]int{3, 3, 3, 3}}, 12},
		{"case-3", args{[]int{3, 4, 3, 2, 5}}, 10},
		{"case-4", args{[]int{1, 2, 3, 4, 5, 6, 5, 4, 3, 2, 1}}, 21},
		{"case-5", args{[]int{5, 4, 3, 2, 1, 2, 3, 4, 5}}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestRectangleArea(tt.args.heights); got != tt.want {
				t.Errorf("largestRectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
