package i40_zui_xiao_de_kge_shu_lcof

import (
	"github.com/TeslaCN/goleetcode/leetcode/i40_zui-xiao-de-kge-shu-lcof/i40_1"
	"github.com/TeslaCN/goleetcode/leetcode/i40_zui-xiao-de-kge-shu-lcof/i40_2"
	"reflect"
	"testing"
)

func Test_1(t *testing.T) {
	testGetLeastNumbers(t, i40_1.GetLeastNumbers)
}

func Test_2(t *testing.T) {
	testGetLeastNumbers(t, i40_2.GetLeastNumbers)
}

func testGetLeastNumbers(t *testing.T, f func([]int, int) []int) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"sample-0", args{[]int{3, 2, 1}, 2}, []int{2, 1}},
		{"sample-1", args{[]int{0, 1, 2, 1}, 1}, []int{0}},
		{"case-0", args{[]int{3, 5, 2, 4, 6, 1, 7}, 7}, []int{7, 6, 5, 4, 3, 2, 1}},
		{"case-1", args{[]int{3, 5, 2, 4, 6, 1, 7}, 1}, []int{1}},
		{"case-2", args{[]int{1}, 1}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.arr, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLeastNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
