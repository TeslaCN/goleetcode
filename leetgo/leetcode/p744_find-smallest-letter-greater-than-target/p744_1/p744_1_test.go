package p744_1

import "testing"

func Test_nextGreatestLetter(t *testing.T) {
	type args struct {
		letters []byte
		target  byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{"case-0", args{[]byte("aabbbcde"), 'b'}, 'c'},
		{"case-1", args{[]byte("aabbbcde"), 'a'}, 'b'},
		{"sample-1", args{[]byte("cfj"), 'c'}, 'f'},
		{"sample-2", args{[]byte("cfj"), 'd'}, 'f'},
		{"sample-3", args{[]byte("cfj"), 'g'}, 'j'},
		{"sample-0", args{[]byte("cfj"), 'a'}, 'c'},
		{"sample-4", args{[]byte("cfj"), 'j'}, 'c'},
		{"sample-5", args{[]byte("cfj"), 'k'}, 'c'},
		{"sample-6", args{[]byte("ab"), 'z'}, 'a'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreatestLetter(tt.args.letters, tt.args.target); got != tt.want {
				t.Errorf("nextGreatestLetter() = %c, want %c", got, tt.want)
			}
		})
	}
}
