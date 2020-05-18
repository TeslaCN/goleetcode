package p152_1

import "testing"

func Test_maxProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-1", args{[]int{-2, 0, -1}}, 0},
		{"sample-0", args{[]int{2, 3, -2, 4}}, 6},
		{"sample-2", args{[]int{-2}}, -2},
		{"case-0", args{[]int{-2, 4, -1, -1}}, 8},
		{"case-1", args{[]int{3, -4, 1, 2, -2, 0, -1}}, 48},
		{"case-2", args{[]int{-1, 2, -3, -4, 0}}, 24},
		{"case-3", args{[]int{-1, 2, -3, 1, -3, 0}}, 18},
		{"case-4", args{[]int{2, -2, 3, 4}}, 12},
		{"case-5", args{[]int{-1}}, -1},
		{"case-6", args{[]int{-1, 2}}, 2},
		{"case-7", args{[]int{0, -1, 0, 2}}, 2},
		{"case-8", args{[]int{1, 3, 1, 2}}, 6},
		{"case-9", args{[]int{1, 3, 0, 0, 0, 1, 2}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
