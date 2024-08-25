package arraystack

import "errors"

type Stack[T any] struct {
	Elements []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Length() int {
	return len(s.Elements)
}

func (s *Stack[T]) IsEmpty() error {
	if s.Length() != 0 {
		return nil
	} else {
		return errors.New("stack is empty")
	}
}

func (s *Stack[T]) Pop() (T, error) {
	var ZeroVal T
	if err := s.IsEmpty(); err != nil {
		return ZeroVal, err
	}
	ZeroVal = s.Elements[s.Length()-1]
	s.Elements = s.Elements[0:s.Length()-1]

	return ZeroVal, nil
}

func (s *Stack[T]) Push(val T) {
	s.Elements = append(s.Elements, val)
}

func (s *Stack[T]) Peek() (T, error) {
	var ZeroVal T
	if err := s.IsEmpty(); err != nil {
		return ZeroVal, err
	}

	return s.Elements[s.Length()-1], nil
}

func (s *Stack[T]) Show() []T {
	return s.Elements
}