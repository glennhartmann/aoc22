package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/glennhartmann/aoclib/stack"
)

var stackRx = regexp.MustCompile(`^\s*\[([A-Z])\]`)

func main() {
	r := bufio.NewReader(os.Stdin)
	reverseStacks := make([]*stack.Stack[string], 0)
outer:
	for {
		s, err := r.ReadString('\n')
		if err == io.EOF {
			panic("file ended short")
		}
		if err != nil {
			panic("unable to read")
		}
		log.Printf("current line: %q", s)

		currCol := 0
		for {
			match := stackRx.FindStringSubmatchIndex(s)
			if len(match) != 4 {
				if strings.HasPrefix(s, " 1") {
					log.Print("end of stack diagram")
					break outer
				}
				log.Print("EOL")
				break
			}
			log.Printf("found match: %v", match)

			currCol = currCol + (match[2]-1)/4 + 1
			id := s[match[2]:match[3]]
			log.Printf("current match is [%s] at col %d", id, currCol)

			s = s[match[1]:]
			log.Printf("remaining input after current match: %q", s)

			m := len(reverseStacks)
			log.Printf("appending %d new columns", currCol-m)
			for i := 0; i < currCol-m; i++ {
				reverseStacks = append(reverseStacks, stack.NewStack[string]())
			}

			log.Printf("pushing %s into rStack[%d]", id, currCol-1)
			reverseStacks[currCol-1].Push(id)
		}
	}

	log.Printf("final column count: %d", len(reverseStacks))
	properStacks := make([]*stack.Stack[string], len(reverseStacks))
	for i := 0; i < len(reverseStacks); i++ {
		sz := reverseStacks[i].Size()
		log.Printf("col %d height: %d", i+1, sz)

		properStacks[i] = stack.NewStackN[string](sz)
		for j := 0; j < sz; j++ {
			v, err := reverseStacks[i].Pop()
			if err != nil {
				panic("bad stack implementation?")
			}
			properStacks[i].Push(v)
		}
	}

	// line with column numbers was already consumed

	// discard the blank line
	if _, err := r.ReadString('\n'); err != nil {
		panic("unable to read")
	}

	for {
		v, err := r.ReadString('\n')
		if err == io.EOF {
			log.Printf("EOF")
			break
		}
		if err != nil {
			panic("invalid input")
		}

		log.Printf("current line: %q", v)
		var count, src, dst int
		n, err := fmt.Sscanf(v, "move %d from %d to %d\n", &count, &src, &dst)
		if n != 3 || err != nil {
			panic("invalid input")
		}

		log.Printf("moving %d from %d to %d", count, src, dst)
		vslice, err := properStacks[src-1].PopN(count)
		if err != nil {
			panic("invalid input")
		}
		properStacks[dst-1].PushN(vslice...)
	}

	var sb strings.Builder
	for _, s := range properStacks {
		v, err := s.Peek()
		if err != nil {
			panic("can't peek")
		}
		sb.WriteString(v)
	}

	fmt.Printf("%s\n", sb.String())
}
