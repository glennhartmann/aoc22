package stackqueuebase

import (
	"fmt"
	"strings"
)

type Base[T any] struct {
	impl []T
	h    SQ[T]
}

func NewBase[T any](h SQ[T]) *Base[T] {
	return NewBaseN[T](0, h)
}

func NewBaseN[T any](size int, h SQ[T]) *Base[T] {
	return &Base[T]{make([]T, 0, size), h}
}

func (b *Base[T]) SetHelper(sq SQ[T]) {
	b.h = sq
}

func (b *Base[T]) Size() int {
	return len(b.impl)
}

func (b *Base[T]) Push(v T) {
	b.impl = append(b.impl, v)
}

func (b *Base[T]) PushN(v ...T) {
	for _, i := range v {
		b.Push(i)
	}
}

func (b *Base[T]) Pop() (T, error) {
	if b.Empty() {
		var r T
		return r, fmt.Errorf("%s is empty", b.h.NameLower())
	}
	v := b.h.Nth(b.impl, 0)
	b.impl = b.h.Rest(b.impl)
	return v, nil
}

func (b *Base[T]) PopN(n int) ([]T, error) {
	if n > b.Size() {
		return nil, fmt.Errorf("can't pop %d elements - there are only %d in the %s", n, b.Size(), b.h.NameLower())
	}
	r := make([]T, 0, n)
	for i := 0; i < n; i++ {
		v, err := b.Pop()
		if err != nil {
			panic("this really shouldn't happen")
		}
		r = append(r, v)
	}
	return r, nil
}

func (b *Base[T]) Empty() bool {
	return len(b.impl) == 0
}

func (b *Base[T]) Peek() (T, error) {
	if b.Empty() {
		var r T
		return r, fmt.Errorf("%s is empty", b.h.NameLower())
	}
	return b.h.Nth(b.impl, 0), nil
}

func (b *Base[T]) PeekN(n int) ([]T, error) {
	if n > b.Size() {
		return nil, fmt.Errorf("can't peek %d elements - there are only %d in the %s", n, b.Size(), b.h.NameLower())
	}

	r := make([]T, 0, n)
	for i := 0; i < n; i++ {
		r = append(r, b.h.Nth(b.impl, i))
	}

	return r, nil
}

func (b *Base[T]) Join(sep string) string {
	if b.Size() == 0 {
		return ""
	}

	ints, err := b.PeekN(b.Size())
	if err != nil {
		panic(fmt.Sprintf("bad %s: Size() = %d, but PeekN(%d) = %+v", b.h.NameLower(), b.Size(), b.Size(), err))
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%d", ints[0]))

	// TODO: add head and tail labels
	for _, i := range ints[1:] {
		sb.WriteString(fmt.Sprintf("%s%d", sep, i))
	}

	return sb.String()
}

type SQ[T any] interface {
	NameLower() string
	Nth([]T, int) T
	Rest([]T) []T
}

// Implements SQ
type Stack[T any] struct{}

func (Stack[T]) NameLower() string     { return "stack" }
func (Stack[T]) Nth(impl []T, n int) T { return impl[len(impl)-n-1] }
func (Stack[T]) Rest(impl []T) []T     { return impl[:len(impl)-1] }

// Implements SQ
type Queue[T any] struct{}

func (Queue[T]) NameLower() string     { return "queue" }
func (Queue[T]) Nth(impl []T, n int) T { return impl[n] }
func (Queue[T]) Rest(impl []T) []T     { return impl[1:] }
