package i62_yuan_quan_zhong_zui_hou_sheng_xia_de_shu_zi_lcof

import (
	"github.com/TeslaCN/goleetcode/leetcode/i62_yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/i62_2"
	"github.com/TeslaCN/goleetcode/leetcode/i62_yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/i62_3"
	"testing"
)

//func Test_1(t *testing.T) {
//	testLastRemaining(t, i62_1.LastRemaining)
//}

//func Test_2(t *testing.T) {
//	testLastRemaining(t, i62_2.LastRemaining)
//}

func Test_3(t *testing.T) {
	testLastRemaining(t, i62_3.LastRemaining)
}

func testLastRemaining(t *testing.T, f func(int, int) int) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample-0", args{5, 3}, 3},
		{"sample-1", args{10, 17}, 2},
		//{"sample-2", args{70866, 116922}, 64165},
		{"case-0", args{1, 1}, 0},
		{"case-1", args{3, 1}, 2},
		{"case-2", args{3, 2}, 2},
		{"case-3", args{1000, 10000}, 3},
		{"case-4", args{10000, 100000}, 8009},
		{"case-5", args{4, 2}, 0},
		{"case-6", args{5, 2}, 2},
		{"case-7", args{6, 2}, 4},
		{"case-8", args{7, 2}, 6},
		{"case-9", args{8, 2}, 0},
		{"case-10", args{4, 3}, 0},
		{"case-11", args{5, 3}, 3},
		{"case-12", args{6, 3}, 0},
		{"case-13", args{7, 3}, 3},
		{"case-14", args{8, 3}, 6},
		{"case-15", args{9, 3}, 0},
		{"case-16", args{3, 3}, 1},
		{"case-max", args{100000, 1000000}, 4337},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("lastRemaining() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMore(t *testing.T) {
	for i := 2; i < 100; i++ {
		t.Logf("%3d %3d %4v", i, 2, i62_2.LastRemaining(i, 2))
	}
	t.Log()
	for i := 2; i < 100; i++ {
		t.Logf("%3d %3d %4v", i, 3, i62_2.LastRemaining(i, 3))
	}
	t.Log()
	for i := 2; i < 100; i++ {
		t.Logf("%3d %3d %4v", i, 4, i62_2.LastRemaining(i, 4))
	}
}
