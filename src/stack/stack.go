package stack

import "fmt"

type Stack[T any] struct {
	impl []T
}

func NewStack[T any]() *Stack[T] {
	return NewStackN[T](0)
}

func NewStackN[T any](size int) *Stack[T] {
	return &Stack[T]{make([]T, 0, size)}
}

func (s *Stack[T]) Size() int {
	return len(s.impl)
}

func (s *Stack[T]) Push(v T) {
	s.impl = append(s.impl, v)
}

func (s *Stack[T]) PushN(v ...T) {
	for _, i := range v {
		s.Push(i)
	}
}

func (s *Stack[T]) Pop() (T, error) {
	if s.Empty() {
		var r T
		return r, fmt.Errorf("stack is empty")
	}
	v := s.impl[len(s.impl)-1]
	s.impl = s.impl[:len(s.impl)-1]
	return v, nil
}

func (s *Stack[T]) PopN(n int) ([]T, error) {
	if n > s.Size() {
		return nil, fmt.Errorf("can't pop %d elements - there are only %d in the stack", n, s.Size())
	}
	r := make([]T, 0, n)
	for i := 0; i < n; i++ {
		v, err := s.Pop()
		if err != nil {
			panic("this really shouldn't happen")
		}
		r = append(r, v)
	}
	return r, nil
}

func (s *Stack[T]) Empty() bool {
	return len(s.impl) == 0
}

func (s *Stack[T]) Peek() (T, error) {
	if s.Empty() {
		var r T
		return r, fmt.Errorf("stack is empty")
	}
	return s.impl[len(s.impl)-1], nil
}

func (s *Stack[T]) PeekN(n int) ([]T, error) {
	if n > s.Size() {
		return nil, fmt.Errorf("can't peek %d elements - there are only %d in the stack", n, s.Size())
	}

	r := make([]T, 0, n)
	for i := 0; i < n; i++ {
		r = append(r, s.impl[len(s.impl)-i-1])
	}

	return r, nil
}
