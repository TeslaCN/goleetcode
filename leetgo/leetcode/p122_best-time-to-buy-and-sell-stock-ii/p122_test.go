package p122_best_time_to_buy_and_sell_stock_ii

import (
	"github.com/TeslaCN/goleetcode/leetcode/p122_best-time-to-buy-and-sell-stock-ii/p122_1"
	"github.com/TeslaCN/goleetcode/leetcode/p122_best-time-to-buy-and-sell-stock-ii/p122_2"
	"testing"
)

func Test_1(t *testing.T) {
	testMaxProfit(t, p122_1.MaxProfit)
}
func Test_2(t *testing.T) {
	testMaxProfit(t, p122_2.MaxProfit)
}
func testMaxProfit(t *testing.T, f func([]int) int) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{[]int{7, 1, 5, 3, 6, 4}}, 7},
		{"sample-1", args{[]int{1, 2, 3, 4, 5}}, 4},
		{"sample-2", args{[]int{7, 6, 4, 3, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
