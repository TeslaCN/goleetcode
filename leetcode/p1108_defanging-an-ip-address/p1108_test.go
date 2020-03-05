package p1108_defanging_an_ip_address

import (
	"github.com/TeslaCN/goleetcode/leetcode/p1108_defanging-an-ip-address/p1108_1"
	"github.com/TeslaCN/goleetcode/leetcode/p1108_defanging-an-ip-address/p1108_2"
	"testing"
)

func Test_1(t *testing.T) {
	testDefangIPaddr(t, p1108_1.DefangIPaddr)
}

func Test_2(t *testing.T) {
	testDefangIPaddr(t, p1108_2.DefangIPaddr)
}

func testDefangIPaddr(t *testing.T, f func(string) string) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"sample-0", args{"1.1.1.1"}, "1[.]1[.]1[.]1"},
		{"sample-1", args{"255.100.50.0"}, "255[.]100[.]50[.]0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.address); got != tt.want {
				t.Errorf("defangIPaddr() = %v, want %v", got, tt.want)
			}
		})
	}
}
