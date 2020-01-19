package p448

import (
	"github.com/TeslaCN/goleetcode/leetcode/p448/p448_1"
	"github.com/TeslaCN/goleetcode/leetcode/p448/p448_2"
	"reflect"
	"testing"
)

type args struct {
	nums []int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{"sample-0", args{[]int{4, 3, 2, 7, 8, 2, 3, 1}}, []int{5, 6}},
	{"case-0", args{[]int{1, 3, 3}}, []int{2}},
}

func Test_1(t *testing.T) {
	testFindDisappearedNumbers(t, p448_1.FindDisappearedNumbers)
}

func Test_2(t *testing.T) {
	testFindDisappearedNumbers(t, p448_2.FindDisappearedNumbers)
}

func testFindDisappearedNumbers(t *testing.T, f func([]int) []int) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindDisappearedNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
