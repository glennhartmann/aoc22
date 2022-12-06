package main

import (
	"fmt"
	"io"
)

func main() {
	var totalPrioritySum int64
	for {
		var s string
		n, err := fmt.Scanln(&s)
		if err == io.EOF {
			break
		}
		if n != 1 || err != nil {
			panic("invalid input")
		}

		countsL := make(rucksack, 52)
		for _, r := range s[:len(s)/2] {
			countsL.inc(r)
		}

		var repeated rune
		for _, r := range s[len(s)/2:] {
			if countsL.get(r) > 0 {
				repeated = r
				break
			}
		}

		totalPrioritySum += int64(getPriority(repeated))
	}
	fmt.Printf("%d\n", totalPrioritySum)
}

type rucksack []int64

func (r rucksack) index(c rune) int {
	return getPriority(c) - 1
}

func (r rucksack) get(c rune) int64 {
	return r[r.index(c)]
}

func (r rucksack) inc(c rune) {
	r[r.index(c)]++
}

func getPriority(r rune) int {
	if int(r) > int('Z') {
		return int(r) - int('a') + 1
	}
	return int(r) - int('A') + 27
}
