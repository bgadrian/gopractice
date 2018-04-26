package __7_intersection

import (
	"container/list"
	"testing"
)

type llint struct {
	a, b   []int
	result int
}

var table = []*llint{
	{
		a:      []int{3, 8, 5, 9, 7, 2, 1},
		b:      []int{4, 6, 7, 2, 1},
		result: 7},
}

func TestSoluton(t *testing.T) {
	for _, ts := range table {
		la, lb := list.New(), list.New()
		for _, a := range ts.a {
			la.PushBack(a)
		}
		for _, b := range ts.b {
			lb.PushBack(b)
		}

		r := solution(la, lb)
		if r == nil || r.Value.(int) != ts.result {
			t.Errorf("exp '%v', got '%v'", ts.result, r)
		}
	}
}
