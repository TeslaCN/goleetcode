package p605_1

import "testing"

func Test_canPlaceFlowers(t *testing.T) {
	type args struct {
		flowerbed []int
		n         int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"sample-0", args{[]int{1, 0, 0, 0, 1}, 1}, true},
		{"sample-1", args{[]int{1, 0, 0, 0, 1}, 2}, false},
		{"sample-2", args{[]int{1}, 0}, true},
		{"sample-3", args{[]int{0}, 1}, true},
		{"case-0", args{[]int{0, 0, 0, 0, 0}, 3}, true},
		{"case-1", args{[]int{0, 0, 1, 0, 0}, 2}, true},
		{"case-2", args{[]int{0, 0, 1, 0, 0}, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPlaceFlowers(tt.args.flowerbed, tt.args.n); got != tt.want {
				t.Errorf("canPlaceFlowers() = %v, want %v", got, tt.want)
			}
		})
	}
}
