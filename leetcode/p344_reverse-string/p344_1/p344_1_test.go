package p344_1

import "testing"

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name   string
		args   args
		expect string
	}{
		{"sample-0", args{[]byte("hello")}, "olleh"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if reverseString(tt.args.s); string(tt.args.s) != tt.expect {
				t.Errorf("Got: [%v], Expect: [%v]", string(tt.args.s), tt.expect)
			}
		})
	}
}
