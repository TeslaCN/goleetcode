package i13_1

import "testing"

func Test_movingCount(t *testing.T) {
	type args struct {
		m int
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{2, 3, 1}, 3},
		{"sample-1", args{3, 1, 0}, 1},
		{"sample-2", args{16, 8, 4}, 15},
		{"sample-3", args{11, 8, 16}, 88},
		{"case-0", args{100, 100, 20}, 6117},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := movingCount(tt.args.m, tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("movingCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
