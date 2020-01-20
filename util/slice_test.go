package util

import "testing"

func BenchmarkConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConvertToIntPointer([]int{0, 1, 2, 3, 4, 5, 6, 7})
	}
}

func TestNums(t *testing.T) {
}
