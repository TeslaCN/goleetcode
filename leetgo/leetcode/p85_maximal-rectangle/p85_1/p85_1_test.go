package p85_1

import "testing"

func Test_maximalRectangle(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{
			[][]byte{
				[]byte("10100"),
				[]byte("10111"),
				[]byte("11111"),
				[]byte("10010"),
			},
		}, 6},
		{"case-0", args{
			[][]byte{
				[]byte("101001"),
				[]byte("101110"),
				[]byte("111111"),
				[]byte("111111"),
				[]byte("101101"),
			},
		}, 12},
		{"case-1", args{
			[][]byte{
				[]byte("10110"),
				[]byte("10111"),
				[]byte("11110"),
				[]byte("10010"),
			},
		}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalRectangle(tt.args.matrix); got != tt.want {
				t.Errorf("maximalRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
