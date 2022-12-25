package main

import (
	"bufio"
	"io"
	"log"
	"math"
	"os"
	"strings"

	"github.com/glennhartmann/aoc22/src/common"
)

func main() {
	lines := make([]string, 0, 50)
	r := bufio.NewReader(os.Stdin)
	for {
		s, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			panic("unable to read")
		}
		lines = append(lines, strings.TrimSpace(s))
	}

	longest := common.Longest(lines)
	var total int64
	for _, s := range lines {
		i := decode(s)
		total += i
		log.Printf("%s ~> %d", common.PadToRight(s, " ", longest), i)
	}
	log.Print("")
	log.Printf("total: %s (%d)", encode(total), total)
}

var low = []string{"0", "1", "2", "1=", "1-"}

func encode(i int64) string {
	j := 1
	for ; iPow(5, j) < i; j++ {
	}
	j--

	b := make([]byte, j+2)
	for x := range b {
		b[x] = '0'
	}

	j++
	for q := int64(1); q > 0; j-- {
		q = i / 5
		r := i % 5
		i = q

		sn := low[r]
		lowDigit := sn[0]
		highDigit := byte('0')
		if len(sn) == 2 {
			lowDigit = sn[1]
			highDigit = sn[0]
		}

		nsn := add(b[j], lowDigit)
		nLowDigit := nsn[0]
		nHighDigit := byte('0')
		if len(nsn) == 2 {
			nLowDigit = nsn[1]
			nHighDigit = nsn[0]
		}

		b[j] = nLowDigit
		nnsn := add(highDigit, nHighDigit)
		if j > 0 {
			b[j-1] = nnsn[0]
		}
	}

	if b[0] == '0' {
		b = b[1:]
	}

	return string(b)
}

func add(n, c byte) string {
	sum := decode(string(n)) + decode(string(c))
	switch sum {
	case -2:
		return "="
	case -1:
		return "-"
	}
	return low[sum]
}

func decode(s string) int64 {
	var total int64
	for i, c := range s {
		placeVal := iPow(5, len(s)-i-1)
		switch c {
		case '0':
		case '1':
			total += placeVal
		case '2':
			total += 2 * placeVal
		case '-':
			total -= placeVal
		case '=':
			total -= 2 * placeVal
		default:
			panic("invalid snafu")
		}
	}
	return total
}

func iPow(b, e int) int64 {
	return int64(math.Pow(float64(b), float64(e)))
}
