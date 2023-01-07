package h

import "errors"

type Stack[T any] struct {
	list   []T
	Length int
}

func (s *Stack[T]) Push(item T) {
	s.list = append(s.list, item)
	s.Length++
}

func (s *Stack[T]) Pop() (*T, error) {
	if s.IsEmpty() {
		return nil, errors.New("nothing to pop")
	}
	el := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]
	s.Length--
	return &el, nil
}

func NewStack[T any](cap int) Stack[T] {
	list := make([]T, cap)
	return Stack[T]{list: list}
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.list) == 0
}
