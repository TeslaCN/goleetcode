package p1108_2

import "strings"

var DefangIPaddr = defangIPaddr

func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}
