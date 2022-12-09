package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/glennhartmann/aoc22/src/common"
)

type cell struct {
	height       int
	visibleInDir []int
}

func newCell(height int) *cell {
	return &cell{height: height, visibleInDir: []int{-1, -1, -1, -1}}
}

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

	compViz(grid)

	final := maxScore(grid)
	log.Printf("final answer: %d", final)
}

type dir int

const (
	left dir = iota
	right
	top
	bottom
)

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

func compVizCellInDir(grid [][]*cell, i, j int, d dir) {
	c := grid[i][j]

	total := 0
	for {
		i, j = moveInDir(dir(d), i, j)
		if i >= len(grid) || i < 0 || j >= len(grid[0]) || j < 0 {
			break
		}
		total++
		if c.height <= grid[i][j].height {
			break
		}
	}
	c.visibleInDir[d] = total
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

func maxScore(grid [][]*cell) int {
	max := -1
	for i := range grid {
		for j := range grid[i] {
			max = common.Max[int](max, scenicScore(grid[i][j]))
		}
	}
	return max
}

func scenicScore(c *cell) int {
	return c.visibleInDir[0] * c.visibleInDir[1] * c.visibleInDir[2] * c.visibleInDir[3]
}
