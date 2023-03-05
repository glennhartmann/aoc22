package main

import (
	"fmt"
	"io"
	"log"
	"strings"
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
	knots := []pair{
		{450, 450},
		{450, 450},
		{450, 450},
		{450, 450},
		{450, 450},
		{450, 450},
		{450, 450},
		{450, 450},
		{450, 450},
	}
	g.set(knots[8].r, knots[8].c)
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

			//log.Printf("(unresolved) h %v, t %v", head, knots[8])
			prev := head
			for i := 0; i < 9; i++ {
				knots[i] = moveTail(prev, knots[i])
				prev = knots[i]
			}
			//log.Printf("(resolved) h %v, t %v", head, knots[8])
			g.set(knots[8].r, knots[8].c)
			//log.Printf("")
			//logGrid(g, head, knots)
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

func logGrid(g *dynamicGrid, head pair, knots []pair) {
	for i := range g.g {
		var sb strings.Builder
		for j := range g.g[i] {
			p := pair{i, j}
			idx := index(knots, p)
			if head == p {
				sb.WriteString("H")
			} else if idx > -1 {
				sb.WriteString(fmt.Sprintf("%d", idx))
			} else {
				sb.WriteString(".")
			}
		}
		log.Print(sb.String())
	}
}

func index(knots []pair, p pair) int {
	for i := range knots {
		if knots[i] == p {
			return i
		}
	}
	return -1
}
