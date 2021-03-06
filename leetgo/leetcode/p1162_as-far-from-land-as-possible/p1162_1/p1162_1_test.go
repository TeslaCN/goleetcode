package p1162_1

import "testing"

func Test_maxDistance(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{[][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}}, 2},
		{"sample-1", args{[][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}}}, 4},
		{"case-0", args{[][]int{{0, 1}}}, 1},
		{"case-1", args{[][]int{{0, 0, 1}}}, 2},
		{"case-2", args{[][]int{{1, 1, 1}}}, -1},
		{"case-3", args{[][]int{{0, 0, 0}}}, -1},
		{"case-4", args{[][]int{{}}}, -1},
		{"case-5", args{[][]int{}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistance(tt.args.grid); got != tt.want {
				t.Errorf("maxDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
