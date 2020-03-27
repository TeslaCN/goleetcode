package p914_1

import "testing"

func Test_hasGroupsSizeX(t *testing.T) {
	type args struct {
		deck []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"sample-0", args{[]int{1, 2, 3, 4, 4, 3, 2, 1}}, true},
		{"sample-1", args{[]int{1, 1, 1, 2, 2, 2, 3, 3}}, false},
		{"sample-2", args{[]int{1}}, false},
		{"sample-3", args{[]int{1, 1}}, true},
		{"sample-4", args{[]int{1, 1, 2, 2, 2, 2}}, true},
		{"sample-5", args{[]int{1, 1, 1, 1, 2, 2, 2, 2, 2, 2}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasGroupsSizeX(tt.args.deck); got != tt.want {
				t.Errorf("hasGroupsSizeX() = %v, want %v", got, tt.want)
			}
		})
	}
}
