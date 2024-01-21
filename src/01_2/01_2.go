package main

import (
	"fmt"

	"github.com/glennhartmann/aoc22/src/01_1/getcalorieses"

	"github.com/glennhartmann/aoclib/heap"
)

func main() {
	calorieses := getcalorieses.Get()
	h := heap.InitN[int64](false, len(calorieses))
	for _, n := range calorieses {
		h.Push(n)
	}
	fmt.Printf("%d\n", h.Pop()+h.Pop()+h.Pop())
}
