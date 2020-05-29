package p207_course_schedule

import (
	"github.com/TeslaCN/goleetcode/leetcode/p207_course-schedule/p207_1"
	"github.com/TeslaCN/goleetcode/leetcode/p207_course-schedule/p207_2"
	"testing"
)

func Test_1(t *testing.T) {
	testCanFinish(t, p207_1.CanFinish)
}

func Test_2(t *testing.T) {
	testCanFinish(t, p207_2.CanFinish)
}

func testCanFinish(t *testing.T, f func(int, [][]int) bool) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case-0", args{3, [][]int{{1, 0}, {1, 2}, {0, 1}}}, false},
		{"sample-0", args{2, [][]int{{1, 0}}}, true},
		{"sample-1", args{2, [][]int{{1, 0}, {0, 1}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}
