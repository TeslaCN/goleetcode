package lcof_46_ba_shu_zi_fan_yi_cheng_zi_fu_chuan_lcof

import (
	"github.com/TeslaCN/goleetcode/leetcode/lcof_46_ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/lcof_46_1"
	"github.com/TeslaCN/goleetcode/leetcode/lcof_46_ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/lcof_46_2"
	"testing"
)

func Test_1(t *testing.T) {
	testTranslateNum(t, lcof_46_1.TranslateNum)
}

func Test_2(t *testing.T) {
	testTranslateNum(t, lcof_46_2.TranslateNum)
}

func testTranslateNum(t *testing.T, f func(int) int) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{25}, 2},
		{"sample-1", args{12258}, 5},
		{"case-0", args{1313}, 4},
		{"case-1", args{1}, 1},
		{"case-2", args{0}, 1},
		{"case-3", args{100000}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.num); got != tt.want {
				t.Errorf("translateNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
