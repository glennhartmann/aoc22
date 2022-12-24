package main

import (
	"fmt"
	"io"
	"log"

	"github.com/glennhartmann/aoc22/src/common"
	dll "github.com/glennhartmann/aoc22/src/doubly_linked_list"
)

const key = 811589153

func main() {
	originalOrder := make([]*dll.Node[int64], 0, 5000)
	nodeList := dll.NewDLL[int64]()
	for {
		var num int64
		n, err := fmt.Scanf("%d\n", &num)
		if err == io.EOF {
			log.Print("EOF")
			break
		}
		if n != 1 || err != nil {
			panic("invalid input")
		}

		nodeList.PushTail(num * key)
		originalOrder = append(originalOrder, nodeList.PeekTailNode())
	}

	//log.Print("Initial arrangement:")
	//log.Print(nodeList.String())
	//log.Print("")

	for mr := 0; mr < 10; mr++ {
		for _, n := range originalOrder {
			if n.Val() == 0 {
				//log.Print("0 does not move:")
				//log.Print(nodeList.String())
				//log.Print()
				continue
			}

			// TODO: refactor into generalized "circular doubly-linked list"?
			np := func(n *dll.Node[int64]) *dll.Node[int64] { return n.Next() }
			ht := func() *dll.Node[int64] { return nodeList.Head() }
			ins := func(nl *dll.DLL[int64], newN, m *dll.Node[int64]) error { return nl.InsertNodeAfter(newN, m) }
			if n.Val() < 0 {
				np = func(n *dll.Node[int64]) *dll.Node[int64] { return n.Prev() }
				ht = func() *dll.Node[int64] { return nodeList.Tail() }
				ins = func(nl *dll.DLL[int64], newN, m *dll.Node[int64]) error { return nl.InsertNodeBefore(newN, m) }
			}

			m := np(n)
			if m == nil {
				m = ht()
			}
			if err := n.RemoveFrom(nodeList); err != nil {
				panic(err)
			}

			avn := common.Abs(n.Val() % nodeList.Len())
			for i := int64(1); i < avn; i++ {
				m = np(m)
				if m == nil {
					m = ht()
				}
			}

			if err := ins(nodeList, n, m); err != nil {
				panic(err)
			}

			prv := n.Prev()
			if prv == nil {
				prv = nodeList.Tail()
			}
			nxt := n.Next()
			if nxt == nil {
				nxt = nodeList.Head()
			}

			//log.Printf("%d moves between %d and %d:", n.Val(), prv.Val(), nxt.Val())
			//log.Print(nodeList.String())
			//log.Print("")
		}

		i0 := int64(0)
		found := false
		for node := nodeList.Head(); node != nil; node = node.Next() {
			if node.Val() == 0 {
				found = true
				break
			}
			i0++
		}

		if !found {
			panic("0 element not found in list")
		}

		i1000 := (i0 + 1000) % nodeList.Len()
		i2000 := (i0 + 2000) % nodeList.Len()
		i3000 := (i0 + 3000) % nodeList.Len()

		e1000, err1000 := nodeList.PeekHeadN(i1000)
		e2000, err2000 := nodeList.PeekHeadN(i2000)
		e3000, err3000 := nodeList.PeekHeadN(i3000)

		if err1000 != nil || err2000 != nil || err3000 != nil {
			panic(fmt.Sprintf("err1000: %v, err2000: %v, err3000: %v", err1000, err2000, err3000))
		}

		log.Printf("coordinate sum after %d rounds of mixing: %d + %d + %d = %d", mr+1, e1000, e2000, e3000, e1000+e2000+e3000)
	}
}
