package p322_coin_change

import (
	"github.com/TeslaCN/goleetcode/leetcode/p322_coin-change/p322_1"
	"github.com/TeslaCN/goleetcode/leetcode/p322_coin-change/p322_2"
	"testing"
)

func Test_1(t *testing.T) {
	testCoinChange(t, p322_1.CoinChange)
}

func Test_2(t *testing.T) {
	testCoinChange(t, p322_2.CoinChange)
}

func testCoinChange(t *testing.T, f func([]int, int) int) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{[]int{1, 2, 5}, 11}, 3},
		{"sample-1", args{[]int{2}, 3}, -1},
		{"case-0", args{[]int{1, 7, 10}, 14}, 2},
		{"case-1", args{[]int{1, 2}, 3}, 2},
		{"case-2", args{[]int{1}, 11}, 11},
		{"case-3", args{[]int{1, 2, 3}, 11}, 4},
		{"case-4", args{[]int{1, 2, 5}, 100}, 20},
		{"case-5", args{[]int{1, 2, 5}, 1000}, 200},
		{"case-6", args{[]int{1}, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
