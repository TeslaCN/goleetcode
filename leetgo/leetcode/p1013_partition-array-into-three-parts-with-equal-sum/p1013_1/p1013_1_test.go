package p1013_1

import "testing"

func Test_canThreePartsEqualSum(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"sample-0", args{[]int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1}}, true},
		{"sample-1", args{[]int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1}}, false},
		{"sample-2", args{[]int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4}}, true},
		{"case-0", args{[]int{0, 0, 0}}, true},
		{"case-1", args{[]int{0, 1, 0, -1, 0}}, true},
		{"case-2", args{[]int{0, 3, 0, -1, -1, -1}}, false},
		{"case-3", args{[]int{-1, -1, -1, 0, 3, 0}}, false},
		{"case-4", args{[]int{0, 0, 0, 1, -1}}, true},
		{"case-5", args{[]int{6, -5, 7, -3, 1}}, false},
		{"case-6", args{[]int{6, -5, -1, 7, -7, 0}}, true},
		{"case-7", args{[]int{6, -5, 1, 7, -3}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canThreePartsEqualSum(tt.args.A); got != tt.want {
				t.Errorf("canThreePartsEqualSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
