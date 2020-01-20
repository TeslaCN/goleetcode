package util

import (
	"fmt"
	"strings"
)

func FormatIntPointerSlice(slice []*int) string {
	if slice == nil || len(slice) < 1 {
		return ""
	}
	s := ""
	for i := 0; i < len(slice); i++ {
		if slice[i] != nil {
			s += fmt.Sprintf("%v ", *slice[i])
		} else {
			s += fmt.Sprintf("%v ", slice[i])
		}
	}
	return strings.TrimSpace(s)
}

func ConvertToIntPointer(nums []int) []*int {
	if nums == nil {
		return nil
	}
	size := len(nums)
	a := make([]*int, size)
	for i := 0; i < size; i++ {
		if nums[i] != 0 {
			a[i] = &nums[i]
		}
	}
	return a
}
