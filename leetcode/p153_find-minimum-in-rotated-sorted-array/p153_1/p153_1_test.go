package p153_1

import "testing"

func Test_findMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{[]int{3, 4, 5, 1, 2}}, 1},
		{"sample-1", args{[]int{4, 5, 6, 7, 0, 1, 2}}, 0},
		{"case-0", args{[]int{1, 2, 3, 4, 5}}, 1},
		{"case-1", args{[]int{2, 3, 4, 5, 1}}, 1},
		{"case-2", args{[]int{1, 2, 3, 4}}, 1},
		{"case-3", args{[]int{1, 2, 3, 4, 0}}, 0},
		{"case-4", args{[]int{4, 5, 6, 0, 1, 2}}, 0},
		{"case-5", args{[]int{4, 5, 0, 1, 2}}, 0},
		{"case-6", args{[]int{6, 0, 1, 2, 3, 4, 5}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
