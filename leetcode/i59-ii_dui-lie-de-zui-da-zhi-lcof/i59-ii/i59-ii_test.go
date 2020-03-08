package i59_ii

import (
	"reflect"
	"testing"
)

func TestMaxQueue(t *testing.T) {
	q := Constructor()
	t.Log(q.Pop_front() == -1)
	t.Log(q.Max_value() == -1)
	q.Push_back(1)
	q.Push_back(2)
	t.Log(q.Max_value() == 2)
	t.Log(q.Pop_front() == 1)
	t.Log(q.Max_value() == 2)
	q.Push_back(1)
	t.Log(q.Pop_front() == 2)
	t.Log(q.Max_value() == 1)
	q.Push_back(3)
	t.Log(q.Max_value() == 3)
	t.Log(q.Pop_front() == 1)
	t.Log(q.Pop_front() == 3)
}

/*
5 4 1 3
*/
func TestMaxQueue_2(t *testing.T) {
	q := Constructor()
	q.Push_back(5)
	q.Push_back(4)
	q.Push_back(1)
	q.Push_back(3)
	t.Log(q.Max_value() == 5)
	t.Log(q.Pop_front() == 5)
	t.Log(q.Max_value() == 4)
	t.Log(q.Pop_front() == 4)
	t.Log(q.Max_value() == 3)
}

func TestMaxQueue_3(t *testing.T) {
	q := Constructor()
	ints := []int{3, 4, 1, 3, 4, 1, 5, 6, 3, 4, 2, 1, 2, 1, 1, 1, 3, 2, 3, 1, 2, 3, 2, 3, 4, 3, 3, 1}
	for _, v := range ints {
		q.Push_back(v)
	}
	t.Log(q.Max_value() == 6)
	var res []int
	for i := 0; i < len(ints); i++ {
		res = append(res, q.Pop_front())
	}
	t.Log(reflect.DeepEqual(ints, res))
}
