package __5_sort_stack_extra_stack

import (
	"testing"
)

var table = [][]int{
	{1, 2, 3, 4, 5},
	{5, 4, 3, 2, 1},
	{1, 6, 3, 2, 5, 4, 2, 1, 8},
}

func TestSolution(t *testing.T) {
	for _, tt := range table {
		s := &stack{}
		for _, v := range tt {
			s.Push(v)
		}
		solution(s)

		//check if it's sorted
		prev, _ := s.Pop()
		for s.IsEmpty() == false {
			curr, _ := s.Pop()
			if curr >= prev {
				prev = curr
				continue
			}
			t.Errorf("for '%v' failed: '%v'",
				tt, s)
			break
		}
	}
}
