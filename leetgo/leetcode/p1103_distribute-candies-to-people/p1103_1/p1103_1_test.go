package p1103_1

import (
	"reflect"
	"testing"
)

func Test_distributeCandies(t *testing.T) {
	type args struct {
		candies    int
		num_people int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"sample-0", args{7, 4}, []int{1, 2, 3, 1}},
		{"sample-1", args{10, 3}, []int{5, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distributeCandies(tt.args.candies, tt.args.num_people); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("distributeCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
