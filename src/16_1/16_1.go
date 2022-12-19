package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var rx = regexp.MustCompile(`Valve ([A-Z]{2}) has flow rate=(\d+); tunnels? leads? to valves? ([A-Z]{2}(, [A-Z]{2})*)`)

type node struct {
	name       string
	flow       int
	open       bool
	neighbours []string
}

type graph map[string]*node

// TODO: go back and actually solve this...
func main() {
	// ./16_1 --graphviz < input.txt | dot -Tpdf -ograph.pdf
	graphviz := len(os.Args) > 1 && os.Args[1] == "--graphviz"

	g := make(graph, 10)
	r := bufio.NewReader(os.Stdin)
	for {
		s, err := r.ReadString('\n')
		if err == io.EOF {
			log.Print("EOF")
			break
		}
		if err != nil {
			panic("unable to read")
		}

		match := rx.FindStringSubmatch(s)
		if len(match) != 5 {
			panic("invalid input")
		}

		src := match[1]
		rate, err := strconv.Atoi(match[2])
		if err != nil {
			panic("invalid int")
		}
		dsts := strings.Split(match[3], ", ")

		log.Printf("src: %q, rate: %d, dsts: %q", src, rate, dsts)

		g[src] = &node{name: src, flow: rate, open: false, neighbours: dsts}
	}

	if graphviz {
		printGraphviz(g)
		os.Exit(0)
	}
}

func printGraphviz(g graph) {
	fmt.Printf("digraph G {\n")

	for name, n := range g {
		fmt.Printf("  %s [label=\"%s\\n%d\"];\n", name, name, n.flow)
		for _, neighbour := range n.neighbours {
			fmt.Printf("  %s -> %s;\n", name, neighbour)
		}
	}

	fmt.Printf("}\n")
}
