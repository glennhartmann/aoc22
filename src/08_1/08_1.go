package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type cell struct {
	height       int
	visible      []state
	highestInDir []int
}

func newCell(height int) *cell {
	return &cell{height: height, visible: make([]state, 4), highestInDir: []int{-1, -1, -1, -1}}
}

type state int

const (
	unknown state = iota
	yes
	no
)

func main() {
	r := bufio.NewReader(os.Stdin)
	grid := make([][]*cell, 0, 50)
	for {
		s, err := r.ReadString('\n')
		if err == io.EOF {
			log.Print("EOF")
			break
		}
		if err != nil {
			panic("unable to read")
		}

		s = strings.TrimSpace(s)
		row := make([]*cell, 0, len(s))
		for _, c := range s {
			i, err := strconv.Atoi(string(c))
			if err != nil {
				panic("invalid input")
			}
			row = append(row, newCell(i))
		}
		grid = append(grid, row)
	}
	log.Print("original input:")
	logGrid(grid, all)

	compViz(grid)
	log.Print("visible from left:")
	logGrid(grid, left)
	log.Print("visible from right:")
	logGrid(grid, right)
	log.Print("visible from top:")
	logGrid(grid, top)
	log.Print("visible from bottom:")
	logGrid(grid, bottom)
	log.Print("visible from any dir:")
	logGrid(grid, any)

	final := countViz(grid)
	log.Printf("final answer: %d", final)
}

type dir int

const (
	left dir = iota
	right
	top
	bottom
	all
	any
)

func logGrid(grid [][]*cell, d dir) {
	for _, row := range grid {
		var sb strings.Builder
		for _, cell := range row {
			if d == all || (d == left && cell.visible[left] == yes) || (d == right && cell.visible[right] == yes) ||
				(d == top && cell.visible[top] == yes) || (d == bottom && cell.visible[bottom] == yes) ||
				(d == any && (cell.visible[left] == yes || cell.visible[right] == yes || cell.visible[top] == yes || cell.visible[bottom] == yes)) {
				sb.WriteString(fmt.Sprintf("%d", cell.height))
			} else {
				sb.WriteString(" ")
			}
		}
		log.Print(sb.String())
	}
}

func compViz(grid [][]*cell) {
	for i := range grid {
		for j := range grid[i] {
			compVizCellInDir(grid, i, j, left)
			compVizCellInDir(grid, i, j, right)
			compVizCellInDir(grid, i, j, top)
			compVizCellInDir(grid, i, j, bottom)
		}
	}
}

func compVizCellInDir(grid [][]*cell, i, j int, d dir) (highest int) {
	c := grid[i][j]

	if c.visible[d] != unknown {
		return c.highestInDir[d]
	}

	k, m := moveInDir(dir(d), i, j)
	if k >= len(grid) || k < 0 || m >= len(grid[0]) || m < 0 {
		c.visible[d] = yes
		c.highestInDir[d] = c.height
		return c.height
	}

	highest = compVizCellInDir(grid, k, m, d)
	if c.height > highest {
		c.highestInDir[d] = c.height
		c.visible[d] = yes
	} else {
		c.highestInDir[d] = highest
		c.visible[d] = no
	}
	return c.highestInDir[d]
}

func moveInDir(d dir, i, j int) (a, b int) {
	switch d {
	case left:
		return i, j - 1
	case right:
		return i, j + 1
	case top:
		return i - 1, j
	case bottom:
		return i + 1, j
	}
	panic("bad direction")
}

func countViz(grid [][]*cell) int {
	total := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j].visible[left] == yes ||
				grid[i][j].visible[right] == yes ||
				grid[i][j].visible[top] == yes ||
				grid[i][j].visible[bottom] == yes {
				total++
			}
		}
	}
	return total
}
