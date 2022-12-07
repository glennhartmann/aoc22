package main

import (
	"fmt"
	"io"
)

func main() {
	var totalPrioritySum int64
outer:
	for {
		triplet := make([]string, 3)
		for i := 0; i < 3; i++ {
			var s string
			n, err := fmt.Scanln(&s)
			if err == io.EOF {
				break outer
			}
			if n != 1 || err != nil {
				panic("invalid input")
			}
			triplet[i] = s
		}

		counts1 := make(rucksack, 52)
		for _, r := range triplet[0] {
			counts1.inc(r)
		}

		counts2 := make(rucksack, 52)
		for _, r := range triplet[1] {
			counts2.inc(r)
		}

		var repeated rune
		for _, r := range triplet[2] {
			if counts1.get(r) > 0 && counts2.get(r) > 0 {
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
