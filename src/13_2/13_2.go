package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/glennhartmann/aoclib/common"
)

const (
	boldStart = "\033[1m"
	boldEnd   = "\033[m"
)

type listlist struct {
	ll  []*listlist // empty != nil
	val int64
}

func (ll *listlist) String() string {
	if ll.ll == nil {
		return fmt.Sprintf("%d", ll.val)
	}
	return fmt.Sprintf("[%s]", common.Fjoin(ll.ll, ",", func(e *listlist) string { return e.String() }))
}

func (left *listlist) cmp(right *listlist) int {
	return left.cmpInternal(right, "")
}

func (left *listlist) cmpInternal(right *listlist, indent string) int {
	//log.Printf("%s- Compare %s vs %s", indent, left.String(), right.String())

	if left.ll == nil && right.ll == nil {
		return left.cmpIntInternal(right, indent+"  ")
	}

	if left.ll != nil && right.ll != nil {
		return left.cmpListInternal(right, indent+"  ")
	}

	return left.cmpMixedInternal(right, indent+"  ")
}

func (left *listlist) cmpIntInternal(right *listlist, indent string) int {
	if left.val == right.val {
		return 0
	}

	if left.val < right.val {
		//log.Printf("%s- Left side is smaller, so inputs are %sin the right order%s", indent, boldStart, boldEnd)
		return -1
	}

	//log.Printf("%s- Right side is smaller, so inputs are %snot%s in the right order", indent, boldStart, boldEnd)
	return 1
}

func (left *listlist) cmpListInternal(right *listlist, indent string) int {
	for i := 0; i < common.Max(len(left.ll), len(right.ll)); i++ {
		if i >= len(left.ll) {
			//log.Printf("%s- Left side ran out of items, so inputs are %sin the right order%s", indent, boldStart, boldEnd)
			return -1
		}
		if i >= len(right.ll) {
			//log.Printf("%s- Right side ran out of items, so inputs are %snot%s in the right order", indent, boldStart, boldEnd)
			return 1
		}

		c := left.ll[i].cmpInternal(right.ll[i], indent+"  ")
		if c == 0 {
			continue
		}
		return c
	}
	return 0
}

func (left *listlist) cmpMixedInternal(right *listlist, indent string) int {
	if left.ll == nil {
		//log.Printf("%s- Mixed types; convert left to [%d] and retry comparison", indent, left.val)
		left.ll = []*listlist{{val: left.val}}
		r := left.cmpInternal(right, indent+"  ")
		left.ll = nil
		return r
	}
	//log.Printf("%s- Mixed types; convert right to [%d] and retry comparison", indent, right.val)
	right.ll = []*listlist{{val: right.val}}
	r := left.cmpInternal(right, indent+"  ")
	right.ll = nil
	return r
}

var (
	divider1 = &listlist{ll: []*listlist{{ll: []*listlist{{val: 2}}}}}
	divider2 = &listlist{ll: []*listlist{{ll: []*listlist{{val: 6}}}}}
)

func main() {
	r := bufio.NewReader(os.Stdin)
	eof := false
	lls := []*listlist{
		divider1,
		divider2,
	}
	for {
		var left, right *listlist
		for i := 0; i < 3; i++ {
			v, err := r.ReadString('\n')
			if err == io.EOF {
				//log.Printf("EOF")
				eof = true
				break
			}
			if err != nil {
				panic("invalid input")
			}
			if i == 2 {
				// expect blank line
				continue
			}
			v = strings.TrimSpace(v)
			if i == 0 {
				left, _ = parseSubLine(v[1:])
			} else {
				right, _ = parseSubLine(v[1:])
			}
		}

		lls = append(lls, left, right)

		//log.Printf("== Pair %d ==", idx)
		// if left.cmp(right) < 0 {
		// 	total += idx
		// }
		//log.Print("")
		if eof {
			break
		}
	}
	//log.Printf("total: %d", total)

	sort.Slice(lls, func(i, j int) bool { return lls[i].cmp(lls[j]) < 0 })
	log.Print(common.Fjoin(lls, "\n", func(e *listlist) string {
		if e.cmp(divider1) == 0 || e.cmp(divider2) == 0 {
			return fmt.Sprintf("%s%s%s", boldStart, e.String(), boldEnd)
		}
		return e.String()
	}))

	total := 1
	for i := range lls {
		if lls[i].cmp(divider1) == 0 || lls[i].cmp(divider2) == 0 {
			total *= i + 1
		}
	}
	log.Printf("total: %d", total)
}

func parseSubLine(s string) (*listlist, int) {
	ii, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return &listlist{val: ii}, len(s)
	}
	ll := &listlist{ll: []*listlist{} /* empty != nil */}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '[':
			ill, il := parseSubLine(s[i+1:])
			i += il
			ll.ll = append(ll.ll, ill)
		case ']':
			return ll, i + 1
		case ',':
			continue
		default:
			idx := strings.IndexAny(s[i:], "[],\n")
			if idx < 0 {
				panic("bad input")
			}
			ii, err := strconv.ParseInt(s[i:i+idx], 10, 64)
			if err != nil {
				panic("bad int")
			}
			ll.ll = append(ll.ll, &listlist{val: ii})
			i += idx - 1
		}
	}
	return ll, len(s)
}
