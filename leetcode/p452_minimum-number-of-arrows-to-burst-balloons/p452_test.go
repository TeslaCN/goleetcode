package p452_minimum_number_of_arrows_to_burst_balloons

import (
	"github.com/TeslaCN/goleetcode/leetcode/p452_minimum-number-of-arrows-to-burst-balloons/p452_1"
	"testing"
)

type args struct {
	points [][]int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"sample-0", args{[][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}}, 2},
	{"sample-1", args{[][]int{{0, 9}, {1, 8}, {7, 8}, {1, 6}, {9, 16}, {7, 13}, {7, 10}, {6, 11}, {6, 9}, {9, 13}}}, 3},
}

func Test_findMinArrowShots(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p452_1.FindMinArrowShots(tt.args.points); got != tt.want {
				t.Errorf("findMinArrowShots() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFindMinArrowShots(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			if got := p452_1.FindMinArrowShots(tt.args.points); got != tt.want {
				b.Logf("findMinArrowShots() = %v, want %v", got, tt.want)
			}
		}
	}
}

func TestCharset(t *testing.T) {
}
