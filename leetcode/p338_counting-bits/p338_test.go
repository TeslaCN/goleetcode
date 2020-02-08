package p338_counting_bits

import (
	"fmt"
	"github.com/TeslaCN/goleetcode/leetcode/p338_counting-bits/p338_1"
	"github.com/TeslaCN/goleetcode/leetcode/p338_counting-bits/p338_2"
	"github.com/TeslaCN/goleetcode/leetcode/p338_counting-bits/p338_3"
	"log"
	"math"
	"reflect"
	"testing"
)

type args struct {
	num int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{"sample-0", args{2}, []int{0, 1, 1}},
	{"sample-1", args{5}, []int{0, 1, 1, 2, 1, 2}},
	//{"case-0", args{100}, []int{0, 1, 1, 2, 1, 2}},
}

func Test_1(t *testing.T) {
	testCountBits(t, p338_1.CountBits)
}
func Test_2(t *testing.T) {
	testCountBits(t, p338_2.CountBits)
}

func Test_3(t *testing.T) {
	testCountBits(t, p338_3.CountBits)
}

func testCountBits(t *testing.T, f func(int) []int) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCount(t *testing.T) {
	for i := 0; i < 100; i++ {
		num := i
		count := 0
		for num > 0 {
			if num&1 == 1 {
				count++
			}
			num >>= 1
		}
		fmt.Printf("%d -> %d\n", i, count)
	}
}

func TestFloat64(t *testing.T) {
	log.Println(math.Log2(35))
	log.Println(math.Log2(32))
	log.Println(math.Pow(2, math.Log2(32)))
	log.Println(int(math.Log2(35)))
	log.Println(math.Log2(35) == float64(int(math.Log2(35))))
	log.Println(math.Log2(32) == float64(int(math.Log2(32))))
	log.Println(math.Log2(0))
	log.Println(math.Log2(1))
	log.Println(math.Log2(1.5))
	log.Println(math.Log2(2))
}
