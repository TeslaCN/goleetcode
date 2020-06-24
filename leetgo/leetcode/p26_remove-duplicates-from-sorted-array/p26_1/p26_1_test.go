package p26_1

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{[]int{1, 1, 2}}, 2},
		{"sample-1", args{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}}, 5},
		{"case-0", args{[]int{1, 2, 3, 4}}, 4},
		{"case-1", args{[]int{1}}, 1},
		{"case-2", args{[]int{11, 11}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
