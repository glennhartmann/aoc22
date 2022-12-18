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
	val := 1
	buf := strings.Builder{}
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
			draw(&buf, val, cycle)
			cycle++
			if (cycle-1)%40 == 0 {
				log.Print(buf.String())
				buf = strings.Builder{}
			}
		}

		val += inc
	}
}

func draw(buf *strings.Builder, spritePos int, cycle int) {
	c := (cycle - 1) % 40
	//log.Printf("(cycle %d) c: %d, s: %d", cycle, c, spritePos)
	if c == spritePos-1 || c == spritePos || c == spritePos+1 {
		buf.WriteString("#")
	} else {
		buf.WriteString(".")
	}
}
