package i08_11_coin_lcci

import (
	i08_11_3 "github.com/TeslaCN/goleetcode/leetcode/i08-11_coin-lcci/i08-11_3"
	"testing"
)

//func Test_1(t *testing.T) {
//	testWaysToChange(t, i08_11_1.WaysToChange)
//}
//
//func Test_2(t *testing.T) {
//	testWaysToChange(t, i08_11_2.WaysToChange)
//}

func Test_3(t *testing.T) {
	testWaysToChange(t, i08_11_3.WaysToChange)
}

func testWaysToChange(t *testing.T, f func(int) int) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{5}, 2},
		{"sample-1", args{10}, 4},
		{"case-0", args{4}, 1},
		{"case-1", args{15}, 6},
		{"case-2", args{20}, 9},
		{"case-3", args{25}, 13},
		{"case-4", args{35}, 24},
		{"case-5", args{40}, 31},
		{"case-6", args{300}, 4464},
		{"case-large", args{1000000}, 332576607},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.n); got != tt.want {
				t.Errorf("waysToChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
