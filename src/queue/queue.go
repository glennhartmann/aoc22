package queue

import "fmt"

// TODO refactor with stack?

type Queue[T any] struct {
	impl []T
}

func NewQueue[T any]() *Queue[T] {
	return NewQueueN[T](0)
}

func NewQueueN[T any](size int) *Queue[T] {
	return &Queue[T]{make([]T, 0, size)}
}

func (s *Queue[T]) Size() int {
	return len(s.impl)
}

func (s *Queue[T]) Push(v T) {
	s.impl = append(s.impl, v)
}

func (s *Queue[T]) PushN(v ...T) {
	for _, i := range v {
		s.Push(i)
	}
}

func (s *Queue[T]) Pop() (T, error) {
	if s.Empty() {
		var r T
		return r, fmt.Errorf("queue is empty")
	}
	v := s.impl[0]
	s.impl = s.impl[1:]
	return v, nil
}

func (s *Queue[T]) PopN(n int) ([]T, error) {
	if n > s.Size() {
		return nil, fmt.Errorf("can't pop %d elements - there are only %d in the queue", n, s.Size())
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

func (s *Queue[T]) Empty() bool {
	return len(s.impl) == 0
}

func (s *Queue[T]) Peek() (T, error) {
	if s.Empty() {
		var r T
		return r, fmt.Errorf("queue is empty")
	}
	return s.impl[0], nil
}

func (s *Queue[T]) PeekN(n int) ([]T, error) {
	if n > s.Size() {
		return nil, fmt.Errorf("can't peek %d elements - there are only %d in the queue", n, s.Size())
	}

	r := make([]T, 0, n)
	for i := 0; i < n; i++ {
		r = append(r, s.impl[i])
	}

	return r, nil
}
