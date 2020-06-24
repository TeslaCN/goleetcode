package p279_perfect_squares

import (
	"github.com/TeslaCN/goleetcode/leetcode/p279_perfect-squares/p279_3"
	"testing"
)

//func Test_1(t *testing.T) {
//	testNumSquares(t, p279_1.NumSquares)
//}

//func Test_2(t *testing.T) {
//	testNumSquares(t, p279_2.NumSquares)
//}

func Test_3(t *testing.T) {
	testNumSquares(t, p279_3.NumSquares)
}

func testNumSquares(t *testing.T, f func(int) int) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{12}, 3},
		{"sample-1", args{13}, 2},
		{"case-0", args{9}, 1},
		{"case-1", args{10}, 2},
		{"case-2", args{101}, 2},
		{"case-3", args{10001}, 2},
		{"case-4", args{10015}, 4},
		{"case-5", args{100000015}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.n); got != tt.want {
				t.Errorf("numSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
