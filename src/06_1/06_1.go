package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	counts := make([]byte, 26)
	numDupes := 0
	buffer := make([]byte, 0, 4)
	for i := 0; ; i++ {
		b, err := r.ReadByte()
		if err != nil {
			panic("???")
		}

		if len(buffer) == 4 {
			c := buffer[0]
			buffer = buffer[1:]
			counts[index(c)]--
			if counts[index(c)] == 1 {
				numDupes--
			}
		}

		buffer = append(buffer, b)
		counts[index(b)]++
		if counts[index(b)] == 2 {
			numDupes++
		}

		if len(buffer) == 4 && numDupes == 0 {
			fmt.Printf("%d\n", i+1)
			return
		}
	}
}

func index(b byte) int {
	return int(b) - int('a')
}
