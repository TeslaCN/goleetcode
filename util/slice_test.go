package util

import (
	"fmt"
	"math"
	"testing"
)

func BenchmarkConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConvertToIntPointer([]int{math.MinInt32, 1, 2, 3, 4, 5, 6, 7})
	}
}

func TestNums(t *testing.T) {
}

func TestAppend_0(t *testing.T) {
	base := make([]int, 0, 1)
	base = append(base, 100)
	fmt.Printf("len(base) = %v, cap(base) = %v\n", len(base), cap(base))
	// equivalent
	// base := []int{100}
	a := append(base, 200)
	b := append(base, 300)
	fmt.Printf("base = %v\na = %v\nb = %v\n", base, a, b)
	fmt.Printf("addr:\nbase = %p\n   a = %p\n   b = %p\n", base, a, b)
	fmt.Printf("len(a) = %v, cap(a) = %v\n", len(a), cap(a))
	fmt.Printf("len(b) = %v, cap(b) = %v\n", len(b), cap(b))
}

func TestAppend_1(t *testing.T) {
	base := make([]int, 0, 2)
	base = append(base, 100)
	fmt.Printf("len(base) = %v, cap(base) = %v\n", len(base), cap(base))
	a := append(base, 200)
	b := append(base, 300)
	fmt.Printf("base = %v\na = %v\nb = %v\n", base, a, b)
	fmt.Printf("addr:\nbase = %p\n   a = %p\n   b = %p\n", base, a, b)
	fmt.Printf("len(a) = %v, cap(a) = %v\n", len(a), cap(a))
	fmt.Printf("len(b) = %v, cap(b) = %v\n", len(b), cap(b))
}
