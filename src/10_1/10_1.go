package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	cycle := 1
	measuringPoints := []int{20, 60, 100, 140, 180, 220}
	mpi := 0
	total := 0
	val := 1
outer:
	for {
		s, err := r.ReadString('\n')
		if err == io.EOF {
			log.Print("EOF")
			break
		}
		if err != nil {
			panic("unable to read")
		}

		inc := 0
		cycles := 0
		if _, after, ok := strings.Cut(s, "addx "); ok {
			after = strings.TrimSpace(after)
			var err error
			inc, err = strconv.Atoi(after)
			if err != nil {
				panic("bad atoi")
			}
			cycles = 2
		} else {
			cycles = 1
		}

		for i := 0; i < cycles; i++ {
			if mpi == len(measuringPoints) {
				break outer
			}
			if cycle == measuringPoints[mpi] {
				log.Printf("cycle %d: val %d, total before %d, total after %d", cycle, val, total, total+val*cycle)
				total += val * cycle
				mpi++
			}
			cycle++
		}

		val += inc
	}
	log.Printf("total: %d", total)
}
