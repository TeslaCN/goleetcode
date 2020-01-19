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
