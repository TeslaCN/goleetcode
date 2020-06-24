package p39_1

import (
	"github.com/TeslaCN/goleetcode/util"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"sample-0", args{[]int{2, 3, 6, 7}, 7}, [][]int{{7}, {2, 2, 3}}},
		{"sample-1", args{[]int{2, 3, 5}, 8}, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{"case-0", args{[]int{2, 3, 6, 7}, 20}, [][]int{{2, 2, 2, 2, 2, 2, 2, 2, 2, 2}, {2, 2, 2, 2, 2, 2, 2, 3, 3}, {2, 2, 2, 2, 2, 2, 2, 6}, {2, 2, 2, 2, 2, 3, 7}, {2, 2, 2, 2, 3, 3, 3, 3}, {2, 2, 2, 2, 3, 3, 6}, {2, 2, 2, 2, 6, 6}, {2, 2, 2, 7, 7}, {2, 2, 3, 3, 3, 7}, {2, 2, 3, 6, 7}, {2, 3, 3, 3, 3, 3, 3}, {2, 3, 3, 3, 3, 6}, {2, 3, 3, 6, 6}, {2, 6, 6, 6}, {3, 3, 7, 7}, {6, 7, 7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum(tt.args.candidates, tt.args.target); !util.EqualsIgnoreOrder(got, tt.want) {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
