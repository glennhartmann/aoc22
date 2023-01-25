package heap

import (
	"container/heap"

	"golang.org/x/exp/constraints"
)

type Heap[T constraints.Ordered] struct {
	hi *heapInternal[T]
}

func Init[T constraints.Ordered](min bool) Heap[T] {
	return InitN[T](min, 0)
}

func InitN[T constraints.Ordered](min bool, size int) Heap[T] {
	h := Heap[T]{&heapInternal[T]{
		impl: make([]T, 0, size),
		min:  min,
	}}
	heap.Init(h.hi)
	return h
}

func (h Heap[T]) Fix(i int)      { heap.Fix(h.hi, i) }
func (h Heap[T]) Pop() T         { return heap.Pop(h.hi).(T) }
func (h Heap[T]) Push(e T)       { heap.Push(h.hi, e) }
func (h Heap[T]) Remove(i int) T { return heap.Remove(h.hi, i).(T) }

type heapInternal[T constraints.Ordered] struct {
	impl []T
	min  bool
}

func (hi heapInternal[T]) Len() int { return len(hi.impl) }
func (hi heapInternal[T]) Less(i, j int) bool {
	if hi.min {
		return hi.impl[i] < hi.impl[j]
	}
	return hi.impl[i] > hi.impl[j]
}
func (hi heapInternal[T]) Swap(i, j int) {
	hi.impl[i], hi.impl[j] = hi.impl[j], hi.impl[i]
}

func (hi *heapInternal[T]) Push(x any) {
	(*hi).impl = append(hi.impl, x.(T))
}

func (hi *heapInternal[T]) Pop() any {
	tmp := hi.impl[len(hi.impl)-1]
	(*hi).impl = hi.impl[:len(hi.impl)-1]
	return tmp
}
