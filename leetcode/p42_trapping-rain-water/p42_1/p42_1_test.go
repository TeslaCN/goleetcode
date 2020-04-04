package p42_1

import "testing"

func Test_trap(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-1", args{[]int{5, 4, 1, 2}}, 1},
		{"sample-0", args{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}}, 6},
		{"case-0", args{[]int{0, 1, 0, 1, 0, 1, 0}}, 2},
		{"case-1", args{[]int{0, 1, 1, 1, 1, 1, 0}}, 0},
		{"case-2", args{[]int{0, 3, 1, 1, 2, 1, 0}}, 2},
		{"case-3", args{[]int{0, 3, 3, 5, 5, 6, 5}}, 0},
		{"case-4", args{[]int{0, 0, 0, 0, 0}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap(tt.args.height); got != tt.want {
				t.Errorf("trap() = %v, want %v", got, tt.want)
			}
		})
	}
}
