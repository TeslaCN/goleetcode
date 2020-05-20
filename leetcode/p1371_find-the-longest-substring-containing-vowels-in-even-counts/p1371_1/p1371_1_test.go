package p1371_1

import "testing"

func Test_findTheLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{"eleetminicoworoep"}, 13},
		{"sample-1", args{"leetcodeisgreat"}, 5},
		{"sample-2", args{"bcbcbc"}, 6},
		{"case-0", args{"abbabbia"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("findTheLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
