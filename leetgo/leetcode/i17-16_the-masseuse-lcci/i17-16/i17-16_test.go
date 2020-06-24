package i17_16

import "testing"

func Test_massage(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{[]int{1, 2, 3, 1}}, 4},
		{"sample-1", args{[]int{2, 7, 9, 3, 1}}, 12},
		{"sample-2", args{[]int{2, 1, 4, 5, 3, 1, 1, 3}}, 12},
		{"case-0", args{[]int{1}}, 1},
		{"case-1", args{[]int{1, 2}}, 2},
		{"case-2", args{[]int{1, 2, 1}}, 2},
		{"case-3", args{[]int{1, 3, 1}}, 3},
		{"case-4", args{[]int{1, 3, 1, 1, 2}}, 5},
		{"case-5", args{[]int{1, 3, 2, 1, 2}}, 5},
		{"case-6", args{[]int{1, 3, 2, 2, 2}}, 5},
		{"case-7", args{[]int{2, 3, 2, 2, 2}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := massage(tt.args.nums); got != tt.want {
				t.Errorf("massage() = %v, want %v", got, tt.want)
			}
		})
	}
}
