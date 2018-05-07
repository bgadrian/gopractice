package __5_sort_stack_extra_stack

import "github.com/pkg/errors"

/* 3.5
Sort Stack: Write a program to sort a stack such that the smallest items are on the top. You can use
an additional temporary stack, but you may not copy the elements into any other data structure
(such as an array) . The stack supports the following operations: push, pop, peek, and isEmpty
*/

type stack struct {
	a []int
}

func (s *stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("empty")
	}
	v := s.a[len(s.a)-1]
	s.a = s.a[:len(s.a)-1]
	return v, nil
}

func (s *stack) Push(v int) {
	s.a = append(s.a, v)
}

func (s *stack) IsEmpty() bool {
	return len(s.a) == 0
}

func (s *stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("empty")
	}
	return s.a[len(s.a)-1], nil
}

func solution(s *stack) {
	b := &stack{}
	var tmp int

	for s.IsEmpty() == false {
		tmp, _ = s.Pop()
		for nv, err := b.Peek(); err != nil && nv > tmp; {
			b.Pop()
			s.Push(nv)
		}
		b.Push(tmp)
	}

	for b.IsEmpty() == false {
		v, _ := b.Pop()
		s.Push(v)
	}
	return
}
