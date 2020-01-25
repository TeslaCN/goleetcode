package p20_1

import (
	"testing"
)

type args struct {
	s string
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{"sample-0", args{"()"}, true},
	{"sample-1", args{"()[]{}"}, true},
	{"sample-2", args{"(]"}, false},
	{"sample-3", args{"([)]"}, false},
	{"sample-4", args{"{[]}"}, true},
	{"sample-5", args{""}, true},
	{"case-0", args{"(["}, false},
	{"case-1", args{")]"}, false},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.args.s); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
