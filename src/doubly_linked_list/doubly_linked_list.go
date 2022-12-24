package doubly_linked_list

import (
	"fmt"
	"strings"
)

// cycles not supported
type DLL[T any] struct {
	head, tail *Node[T]
	length     int64
}

func NewDLL[T any]() *DLL[T] { return &DLL[T]{} }

func (d *DLL[T]) String() string {
	var sb strings.Builder
	sb.WriteString("[")

	first := true
	for n := d.head; n != nil; n = n.next {
		if !first {
			sb.WriteString(" <=> ")
		}
		first = false
		sb.WriteString(fmt.Sprintf("%v", n.val))
	}

	sb.WriteString(fmt.Sprintf("] (%d items)", d.length))
	return sb.String()
}

func (d *DLL[T]) Head() *Node[T] { return d.head }
func (d *DLL[T]) Tail() *Node[T] { return d.tail }
func (d *DLL[T]) Len() int64     { return d.length }

func (d *DLL[T]) PushHead(val T) {
	n := NewNode(val)
	n.next = d.head

	if d.head != nil {
		d.head.prev = n
	}
	d.head = n

	if d.tail == nil {
		d.tail = n
	}

	d.length++
}

func (d *DLL[T]) PushTail(val T) {
	n := NewNode(val)
	n.prev = d.tail

	if d.tail != nil {
		d.tail.next = n
	}
	d.tail = n

	if d.head == nil {
		d.head = n
	}

	d.length++
}

func (d *DLL[T]) PopHead() (T, error) {
	v, err := d.PeekHead()
	if err != nil {
		return v, fmt.Errorf("d.PeekHead(): %w", err)
	}
	d.head = d.head.next

	d.length--

	return v, nil
}

func (d *DLL[T]) PopTail() (T, error) {
	v, err := d.PeekTail()
	if err != nil {
		return v, fmt.Errorf("d.PeekTail(): %w", err)
	}
	d.tail = d.tail.next

	d.length--

	return v, nil
}

func (d *DLL[T]) PeekHead() (T, error) {
	if d.head == nil {
		var r T
		return r, fmt.Errorf("d.head == nil")
	}
	return d.head.val, nil
}

func (d *DLL[T]) PeekHeadN(n int64) (T, error) {
	if n >= d.length {
		var r T
		return r, fmt.Errorf("wanted item %d; list only contains %d items", n, d.length)
	}

	// TODO: could be optimized by going backwards if it's closer to the end
	node := d.head
	for i := int64(0); i < n; i++ {
		if node == nil {
			panic("ran off the end of the list")
		}
		node = node.next
	}

	return node.val, nil
}

func (d *DLL[T]) PeekTail() (T, error) {
	if d.tail == nil {
		var r T
		return r, fmt.Errorf("d.tail == nil")
	}
	return d.tail.val, nil
}

func (d *DLL[T]) PeekTailN(n int64) (T, error) {
	if n >= d.length {
		var r T
		return r, fmt.Errorf("wanted item %d; list only contains %d items", n, d.length)
	}

	// TODO: could be optimized by going forwards if it's closer to the start
	node := d.tail
	for i := int64(0); i < n; i++ {
		if node == nil {
			panic("ran off the start of the list")
		}
		node = node.prev
	}

	return node.val, nil
}

func (d *DLL[T]) PeekHeadNode() *Node[T] { return d.head }
func (d *DLL[T]) PeekTailNode() *Node[T] { return d.tail }

func (d *DLL[T]) InsertAfter(val T, n *Node[T]) error {
	return d.InsertNodeAfter(NewNode(val), n)
}

func (d *DLL[T]) InsertNodeAfter(newN, n *Node[T]) error {
	if n.next == nil && d.tail != n {
		return fmt.Errorf("n.next == nil && d.tail != n -- n is not part of this list?")
	}

	newN.next = n.next
	newN.prev = n
	n.next = newN

	if newN.next == nil {
		d.tail = newN
	} else {
		newN.next.prev = newN
	}

	d.length++

	return nil
}

func (d *DLL[T]) InsertBefore(val T, n *Node[T]) error {
	return d.InsertNodeBefore(NewNode(val), n)
}

func (d *DLL[T]) InsertNodeBefore(newN, n *Node[T]) error {
	if n.prev == nil && d.head != n {
		return fmt.Errorf("n.prev == nil && d.head != n -- n is not part of this list?")
	}

	newN.prev = n.prev
	newN.next = n
	n.prev = newN

	if newN.prev == nil {
		d.head = newN
	} else {
		newN.prev.next = newN
	}

	d.length++

	return nil
}

type Node[T any] struct {
	next, prev *Node[T]
	val        T
}

func NewNode[T any](val T) *Node[T] { return &Node[T]{val: val} }

func (n *Node[T]) Val() T         { return n.val }
func (n *Node[T]) Next() *Node[T] { return n.next }
func (n *Node[T]) Prev() *Node[T] { return n.prev }

func (n *Node[T]) RemoveFrom(dll *DLL[T]) error {
	if n.prev == nil && n != dll.head {
		return fmt.Errorf("n.prev == nil && n != dll.head -- n is not part of this list?")
	}
	if n.next == nil && n != dll.tail {
		return fmt.Errorf("n.next == nil && n != dll.tail -- n is not part of this list?")
	}

	if n == dll.head {
		dll.head = n.next
	}
	if n == dll.tail {
		dll.tail = n.prev
	}

	if n.next != nil {
		n.next.prev = n.prev
	}
	if n.prev != nil {
		n.prev.next = n.next
	}

	dll.length--

	return nil
}

// TODO: do these two need refactoring?
func (n *Node[T]) HeadIndexIn(dll *DLL[T]) (int64, error) {
	i := int64(0)
	for node := dll.head; node != nil; node = node.next {
		if n == node {
			return i, nil
		}
		i++
	}
	return -1, fmt.Errorf("node not found in list")
}

func (n *Node[T]) TailIndexIn(dll *DLL[T]) (int64, error) {
	i := int64(0)
	for node := dll.tail; node != nil; node = node.prev {
		if n == node {
			return i, nil
		}
		i++
	}
	return -1, fmt.Errorf("node not found in list")
}
