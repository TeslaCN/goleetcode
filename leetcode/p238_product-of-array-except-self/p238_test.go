package p238_product_of_array_except_self

import (
	"github.com/TeslaCN/goleetcode/leetcode/p238_product-of-array-except-self/p238_0"
	"github.com/TeslaCN/goleetcode/leetcode/p238_product-of-array-except-self/p238_1"
	"reflect"
	"testing"
)

func Test_0(t *testing.T) {
	testProductExceptSelf(t, p238_0.ProductExceptSelf)
}

func Test_1(t *testing.T) {
	testProductExceptSelf(t, p238_1.ProductExceptSelf)
}
func testProductExceptSelf(t *testing.T, f func([]int) []int) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"sample-0", args{[]int{1, 2, 3, 4}}, []int{24, 12, 8, 6}},
		{"case-0", args{[]int{0, 1, 2, 3, 4}}, []int{24, 0, 0, 0, 0}},
		{"case-1", args{[]int{0, 0, 1, 2, 3, 4}}, []int{0, 0, 0, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productExceptSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}
