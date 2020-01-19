package p461_1

import "testing"

type args struct {
	x int
	y int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"sample-0", args{1, 4}, 2},
	{"case-0", args{1, 6}, 3},
	{"case-1", args{9, 6}, 4},
	{"case-2", args{-145143, -313432}, 12},
}

func TestHammingDistance(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HammingDistance(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("HammingDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
