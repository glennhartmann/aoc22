package main

import (
	"fmt"
	"io"
	"log"
)

func main() {
	grid := make([][][]bool, 20)
	for x := range grid {
		grid[x] = make([][]bool, 20)
		for y := range grid[x] {
			grid[x][y] = make([]bool, 20)
		}
	}
	for {
		var x, y, z int64
		n, err := fmt.Scanf("%d,%d,%d\n", &x, &y, &z)
		if err == io.EOF {
			break
		}
		if n != 3 || err != nil {
			panic("invalid input")
		}

		grid[x][y][z] = true
	}

	area := 0
	for x := range grid {
		for y := range grid[x] {
			for z := range grid[x][y] {
				if !grid[x][y][z] {
					continue
				}
				a := 6
				if x < len(grid)-1 && grid[x+1][y][z] {
					a--
				}
				if x > 0 && grid[x-1][y][z] {
					a--
				}
				if y < len(grid[x])-1 && grid[x][y+1][z] {
					a--
				}
				if y > 0 && grid[x][y-1][z] {
					a--
				}
				if z < len(grid[x][y])-1 && grid[x][y][z+1] {
					a--
				}
				if z > 0 && grid[x][y][z-1] {
					a--
				}
				area += a
			}
		}
	}

	log.Printf("area: %d", area)
}
