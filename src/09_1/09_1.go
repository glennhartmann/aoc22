package main

import (
	"fmt"
	"io"
	"log"
)

type dynamicGrid struct {
	g [][]bool
}

func newDG() *dynamicGrid {
	g := make([][]bool, 500)
	for i := range g {
		g[i] = make([]bool, 500)
	}
	return &dynamicGrid{g: g}
}

func (dg *dynamicGrid) get(r, c int) bool {
	dg.maybeGrow(r, c)
	return dg.g[r][c]
}

func (dg *dynamicGrid) set(r, c int) {
	dg.maybeGrow(r, c)
	dg.g[r][c] = true
}

func (dg *dynamicGrid) maybeGrow(r, c int) {
	growRow := r >= len(dg.g)
	if growRow {
		dg.g = append(dg.g, make([][]bool, r-len(dg.g)+20)...)
	}
	for i := range dg.g {
		minSize := c - len(dg.g[i]) + 20
		if minSize > 0 {
			dg.g[i] = append(dg.g[i], make([]bool, minSize)...)
		}
	}
}

type pair struct{ r, c int }

type dir int

const (
	left dir = iota
	right
	up
	down
)

func dirFromChar(c byte) dir {
	switch c {
	case 'L':
		return left
	case 'R':
		return right
	case 'U':
		return up
	case 'D':
		return down
	}
	panic("bad dir")
}

func main() {
	g := newDG()
	head := pair{450, 450}
	tail := pair{450, 450}
	g.set(tail.r, tail.c)
	for {
		var dir byte
		var dist int
		n, err := fmt.Scanf("%c %d\n", &dir, &dist)
		if err == io.EOF {
			break
		}
		if n != 2 || err != nil {
			panic("invalid input")
		}

		log.Printf("%c %d", dir, dist)
		for i := 0; i < dist; i++ {
			d := dirFromChar(dir)
			head.r, head.c = moveInDir(d, head.r, head.c)

			log.Printf("(unresolved) h %v, t %v", head, tail)
			tail = moveTail(head, tail)
			log.Printf("(resolved) h %v, t %v", head, tail)
			g.set(tail.r, tail.c)
		}
	}

	total := 0
	for i := range g.g {
		for j := range g.g[i] {
			if g.get(i, j) {
				total++
			}
		}
	}
	log.Printf("final: %d", total)
}

func moveInDir(d dir, i, j int) (a, b int) {
	switch d {
	case left:
		return i, j - 1
	case right:
		return i, j + 1
	case up:
		return i - 1, j
	case down:
		return i + 1, j
	}
	panic("bad direction")
}

func moveTail(head, tail pair) pair {
	if adjacent(head, tail) {
		return tail
	}
	rDiff := head.r - tail.r
	cDiff := head.c - tail.c
	tail.r += sign(rDiff)
	tail.c += sign(cDiff)
	return tail
}

func sign(x int) int {
	if x == 0 {
		return 0
	}
	if x < 0 {
		return -1
	}
	return 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func adjacent(head, tail pair) bool {
	return abs(head.c-tail.c) <= 1 && abs(head.r-tail.r) <= 1
}
