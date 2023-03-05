// Package stack implements a fully generic stack. It is essentially a public
// interface for internal/stackequeuebase.
package stack

import (
	"github.com/glennhartmann/aoc22/src/internal/stackqueuebase"
)

// Stack is a generic queue.
type Stack[T any] struct {
	*stackqueuebase.Base[T]
}

// NewStack creates an empty Stack.
func NewStack[T any]() *Stack[T] {
	return NewStackN[T](0)
}

// NewStackN creates an empty Stack with |size| elements worth of preallocated
// memory.
func NewStackN[T any](size int) *Stack[T] {
	return &Stack[T]{stackqueuebase.NewBaseN[T](size, stackqueuebase.Stack[T]{})}
}
