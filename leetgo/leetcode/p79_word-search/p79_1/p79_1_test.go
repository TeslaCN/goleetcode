package p79_1

import "testing"

func Test_exist(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"sample-0", args{[][]byte{
			[]byte("ABCE"),
			[]byte("SFCS"),
			[]byte("ADEE"),
		}, "ABCCED"}, true},
		{"sample-1", args{[][]byte{
			[]byte("ABCE"),
			[]byte("SFCS"),
			[]byte("ADEE"),
		}, "SEE"}, true},
		{"sample-2", args{[][]byte{
			[]byte("ABCE"),
			[]byte("SFCS"),
			[]byte("ADEE"),
		}, "ABCB"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exist(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
