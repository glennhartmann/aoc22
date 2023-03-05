// Package stackqueuebase implements core functionality fully generic stacks
// and queues.
package stackqueuebase

import (
	"fmt"
	"strings"
)

// Base implements stack and queue functionality, given the appropriate |impl|
// helper.
type Base[T any] struct {
	impl []T
	h    SQ[T]
}

// NewBase creates a new Base object.
func NewBase[T any](h SQ[T]) *Base[T] {
	return NewBaseN[T](0, h)
}

// NewBaseN creates a new Base object with |size| elements worth of
// preallocated memory.
func NewBaseN[T any](size int, h SQ[T]) *Base[T] {
	return &Base[T]{make([]T, 0, size), h}
}

// Size returns the number of elements in the stack or queue.
func (b *Base[T]) Size() int {
	return len(b.impl)
}

// Push pushes an item into the stack or queue.
func (b *Base[T]) Push(v T) {
	b.impl = append(b.impl, v)
}

// PushN pushes N items into the stack or queue.
func (b *Base[T]) PushN(v ...T) {
	for _, i := range v {
		b.Push(i)
	}
}

// Pop pops an item from the stack or queue. If the stack or queue is empty, an
// error is returned instead and the stack or queue remains unmodified.
func (b *Base[T]) Pop() (T, error) {
	if b.Empty() {
		var r T
		return r, fmt.Errorf("%s is empty", b.h.NameLower())
	}
	v := b.h.Nth(b.impl, 0)
	b.impl = b.h.Rest(b.impl)
	return v, nil
}

// Pop pops N items from the stack or queue. If the stack or queue has fewer
// than N elements, an error is returned instead and the stack or queue remains
// unmodified.
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

// Empty returns whether or not the stack or queue is empty.
func (b *Base[T]) Empty() bool {
	return len(b.impl) == 0
}

// Pop returns an item from the stack or queue without removing it. If the
// stack or queue is empty, an error is returned instead.
func (b *Base[T]) Peek() (T, error) {
	if b.Empty() {
		var r T
		return r, fmt.Errorf("%s is empty", b.h.NameLower())
	}
	return b.h.Nth(b.impl, 0), nil
}

// Pop returns N items from the stack or queue without removing them. If the
// stack or queue has fewer than N elements, an error is returned instead.
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

// Join creates a string by combining each element from the stack or queue,
// with |sep| between each pair.
func (b *Base[T]) Join(sep string) string {
	if b.Size() == 0 {
		return ""
	}

	vals, err := b.PeekN(b.Size())
	if err != nil {
		panic(fmt.Sprintf("bad %s: Size() = %d, but PeekN(%d) = %+v", b.h.NameLower(), b.Size(), b.Size(), err))
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%v", vals[0]))

	// TODO: add head and tail labels
	for _, v := range vals[1:] {
		sb.WriteString(fmt.Sprintf("%s%v", sep, v))
	}

	return sb.String()
}

// SQ is an interface for the helper object that implements either
// stack-specific or queue-specific functionality.
type SQ[T any] interface {
	// NameLower returns the name of this object ("stack" or "queue").
	NameLower() string

	// Nth returns the Nth item from the slice.
	Nth([]T, int) T

	// Rest returns all but the first item from the slice.
	Rest([]T) []T
}

// Stack implements SQ for a stack.
type Stack[T any] struct{}

// NameLower returns the name of this object ("stack").
func (Stack[T]) NameLower() string { return "stack" }

// Nth returns the Nth item from the slice.
func (Stack[T]) Nth(impl []T, n int) T { return impl[len(impl)-n-1] }

// Rest returns all but the first item from the slice.
func (Stack[T]) Rest(impl []T) []T { return impl[:len(impl)-1] }

// Queue implements SQ for a queue.
type Queue[T any] struct{}

// NameLower returns the name of this object ("queue").
func (Queue[T]) NameLower() string { return "queue" }

// Nth returns the Nth item from the slice.
func (Queue[T]) Nth(impl []T, n int) T { return impl[n] }

// Rest returns all but the first item from the slice.
func (Queue[T]) Rest(impl []T) []T { return impl[1:] }
