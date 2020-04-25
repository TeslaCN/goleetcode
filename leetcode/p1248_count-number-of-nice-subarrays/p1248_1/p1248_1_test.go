package p1248_1

import "testing"

func Test_numberOfSubarrays(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-2", args{[]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2}, 16},
		{"sample-0", args{[]int{1, 1, 2, 1, 1}, 3}, 2},
		{"sample-1", args{[]int{2, 4, 6}, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfSubarrays(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("numberOfSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
