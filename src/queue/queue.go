package queue

import (
	"github.com/glennhartmann/aoc22/src/internal/stackqueuebase"
)

type Queue[T any] struct {
	*stackqueuebase.Base[T]
}

func NewQueue[T any]() *Queue[T] {
	return NewQueueN[T](0)
}

func NewQueueN[T any](size int) *Queue[T] {
	return &Queue[T]{stackqueuebase.NewBaseN[T](size, stackqueuebase.Queue[T]{})}
}
