package p210_1

import (
	"reflect"
	"testing"
)

func Test_findOrder(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name  string
		args  args
		wants [][]int
	}{
		{"sample-0", args{2, [][]int{{1, 0}}}, [][]int{{0, 1}}},
		{"sample-1", args{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}}, [][]int{{0, 1, 2, 3}, {0, 2, 1, 3}}},
		{"sample-2", args{2, [][]int{{0, 1}, {1, 0}}}, [][]int{{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, hit := findOrder(tt.args.numCourses, tt.args.prerequisites), false
			for i := 0; !hit && i < len(tt.wants); i++ {
				hit = reflect.DeepEqual(got, tt.wants[i])
			}
			if !hit {
				t.Errorf("findOrder() = %v, want %v", got, tt.wants)
			}
		})
	}
}
