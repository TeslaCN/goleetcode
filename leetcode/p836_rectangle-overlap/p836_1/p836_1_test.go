package p836_1

import "testing"

func Test_isRectangleOverlap(t *testing.T) {
	type args struct {
		rec1 []int
		rec2 []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"sample-0", args{[]int{0, 0, 2, 2}, []int{1, 1, 3, 3}}, true},
		{"sample-1", args{[]int{0, 0, 1, 1}, []int{1, 0, 2, 1}}, false},
		{"case-0", args{[]int{0, 0, 2, 2}, []int{-1, 1, 1, 3}}, true},
		{"case-1", args{[]int{0, 0, 2, 2}, []int{1, -1, 3, 1}}, true},
		{"case-2", args{[]int{0, 0, 2, 2}, []int{1, 1, 3, 2}}, true},
		{"case-3", args{[]int{0, 0, 2, 2}, []int{2, 0, 3, 1}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRectangleOverlap(tt.args.rec1, tt.args.rec2); got != tt.want {
				t.Errorf("isRectangleOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}
