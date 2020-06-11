package util

import (
	"sort"
	"strconv"
)

func EqualsIgnoreOrder(actual, expected [][]int) bool {
	if len(expected) != len(actual) {
		return false
	}
	var expectedKeys []string
	var actualKeys []string
	for _, values := range expected {
		key := ""
		for _, value := range values {
			key += strconv.Itoa(value) + " "
		}
		expectedKeys = append(expectedKeys, key)
	}
	for _, values := range actual {
		key := ""
		for _, value := range values {
			key += strconv.Itoa(value) + " "
		}
		actualKeys = append(actualKeys, key)
	}
	sort.Strings(expectedKeys)
	//log.Printf("expect: %v", expectedKeys)
	sort.Strings(actualKeys)
	//log.Printf("actual: %v", actualKeys)
	for i, expect := range expectedKeys {
		if actualKeys[i] != expect {
			return false
		}
	}
	return true
}

func EqualsIgnoreOrderString(actual, expected [][]string) bool {
	if len(actual) != len(expected) {
		return false
	}
	less := func(i, j int) bool {
		m, n := len(actual[i]), len(actual[j])
		if m == 0 || n == 0 {
			return n > 0
		}
		if m < n {
			return true
		} else if m > n {
			return false
		}
		for a := 0; a < m; a++ {
			if actual[i][a] != actual[j][a] {
				return actual[i][a] < actual[j][a]
			}
		}
		return true
	}
	sort.Slice(actual, less)
	sort.Slice(expected, less)
	//return reflect.DeepEqual(actual, expected)
	panic("Unsupported")
}
