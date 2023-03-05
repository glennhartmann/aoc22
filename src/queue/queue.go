// Package queue implements a fully generic queue. It is essentially a public
// interface for internal/stackequeuebase.
package queue

import (
	"github.com/glennhartmann/aoc22/src/internal/stackqueuebase"
)

// Queue is a generic queue.
type Queue[T any] struct {
	*stackqueuebase.Base[T]
}

// NewQueue creates an empty Queue.
func NewQueue[T any]() *Queue[T] {
	return NewQueueN[T](0)
}

// NewQueueN creates an empty Queue with |size| elements worth of preallocated
// memory.
func NewQueueN[T any](size int) *Queue[T] {
	return &Queue[T]{stackqueuebase.NewBaseN[T](size, stackqueuebase.Queue[T]{})}
}
