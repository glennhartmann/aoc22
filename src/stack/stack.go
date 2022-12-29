package stack

import (
	"github.com/glennhartmann/aoc22/src/internal/stackqueuebase"
)

type Stack[T any] struct {
	*stackqueuebase.Base[T]
}

func NewStack[T any]() *Stack[T] {
	return NewStackN[T](0)
}

func NewStackN[T any](size int) *Stack[T] {
	return &Stack[T]{stackqueuebase.NewBaseN[T](size, stackqueuebase.Stack[T]{})}
}
