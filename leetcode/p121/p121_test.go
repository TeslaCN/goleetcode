package p121

import (
	"github.com/TeslaCN/goleetcode/leetcode/p121/p121_1"
	"github.com/TeslaCN/goleetcode/leetcode/p121/p121_2"
	"log"
	"math/rand"
	"testing"
)

type args struct {
	prices []int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"sample-0", args{[]int{7, 1, 5, 3, 6, 4}}, 5},
	{"sample-1", args{[]int{7, 6, 4, 3, 1}}, 0},
	{"case-0", args{[]int{7, 1, 5, 1, 3, 6, 4, 6}}, 5},
	{"case-1", args{[]int{7, 1, 8, 5, 1, 3, 6, 7, 4, 6}}, 7},
	{"case-2", args{[]int{}}, 0},
	{"case-3", args{[]int{1}}, 0},
	{"case-4", args{[]int{2, 4, 1}}, 2},
}

// Add massive data
func init() {
	maxValue := 2000000000
	ints := make([]int, 100000)
	ints[0] = 0
	ints[len(ints)-1] = maxValue
	for i := 1; i < len(ints)-1; i++ {
		ints[i] = rand.Intn(maxValue)
	}

	tests = append(tests, struct {
		name string
		args args
		want int
	}{name: "massive", args: args{ints}, want: maxValue})
	log.Println("Init test case.")
}

func Test_1(t *testing.T) {
	testMaxProfit(t, p121_1.MaxProfit)
}

func Test_2(t *testing.T) {
	testMaxProfit(t, p121_2.MaxProfit)
}

func testMaxProfit(t *testing.T, f func([]int) int) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.prices); got != tt.want {
				t.Errorf("MaxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
