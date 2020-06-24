package p365_1

import "testing"

func Test_canMeasureWater(t *testing.T) {
	type args struct {
		x int
		y int
		z int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"sample-5", args{4, 6, 8}, true},
		{"sample-0", args{3, 5, 4}, true},
		{"sample-1", args{2, 6, 5}, false},
		{"sample-2", args{0, 0, 5}, false},
		{"sample-3", args{1, 1, 12}, false},
		{"sample-4", args{13, 11, 1}, true},
		{"case-0", args{1, 2, 3}, true},
		{"case-1", args{2, 4, 5}, false},
		{"case-2", args{2, 4, 50001}, false},
		{"case-3", args{0, 5, 50000}, false},
		{"case-4", args{19, 13, 12}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canMeasureWater(tt.args.x, tt.args.y, tt.args.z); got != tt.want {
				t.Errorf("canMeasureWater() = %v, want %v", got, tt.want)
			}
		})
	}
}
