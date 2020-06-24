package p406_queue_reconstruction_by_height

import (
	"github.com/TeslaCN/goleetcode/leetcode/p406_queue-reconstruction-by-height/p406_1"
	"github.com/TeslaCN/goleetcode/leetcode/p406_queue-reconstruction-by-height/p406_2"
	"reflect"
	"testing"
)

type args struct {
	people [][]int
}

func Test_1(t *testing.T) {
	testReconstructQueue(t, p406_1.ReconstructQueue)
}

func Test_2(t *testing.T) {
	testReconstructQueue(t, p406_2.ReconstructQueue)
}

func testReconstructQueue(t *testing.T, f func([][]int) [][]int) {
	var tests = []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "sample-0",
			args: args{[][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}},
			want: [][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.people); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reconstructQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}
