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
