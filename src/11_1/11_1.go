package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"

	"github.com/glennhartmann/aoclib/queue"
)

type readerState int

const (
	blank readerState = iota
	monkeyLabel
	startingItems
	operation
	test
	trueCondition
	falseCondition

	numStates
)

const logVerbose = true

var (
	startingItemsRx = regexp.MustCompile(`\s*Starting items: (\d+(, \d+)*)`)
	operationRx     = regexp.MustCompile(`\s*Operation: new = (\w+) (.) (\w+)`)
	testRx          = regexp.MustCompile(`\s*Test: divisible by (\d+)`)
	conditionRx     = regexp.MustCompile(`\s*If (true|false): throw to monkey (\d+)`)
)

type monkey struct {
	index              int
	items              *queue.Queue[int]
	operation          op
	testDivisor        int
	trueConditionDest  int
	falseConditionDest int
	inspectCount       int
}

type op struct {
	right    operand
	operator byte
}

type operand struct {
	old bool
	num int
}

func main() {
	first := true
	monkeys := make([]*monkey, 0, 7)
	r := bufio.NewReader(os.Stdin)
outer:
	for {
		m := &monkey{index: len(monkeys)}
		log.Printf("monkey %d:", m.index)
		for i := readerState(0); i < numStates; i++ {
			if first {
				// first money doesn't have a preceding blank line
				first = false
				continue
			}

			s, err := r.ReadString('\n')
			if err == io.EOF {
				log.Printf("EOF - monkey %d discarded", m.index)
				break outer
			}
			if err != nil {
				panic("unable to read")
			}

			switch i {
			case blank, monkeyLabel:
				continue
			case startingItems:
				match := startingItemsRx.FindStringSubmatch(s)
				if len(match) != 3 {
					panic("invalid input")
				}
				itemsStr := strings.Split(match[1], ", ")
				log.Printf("  starting items: %v", itemsStr)

				m.items = queue.NewQueueN[int](len(itemsStr))
				for _, s := range itemsStr {
					item, err := strconv.Atoi(s)
					if err != nil {
						panic("invalid int")
					}
					m.items.Push(item)
				}
			case operation:
				match := operationRx.FindStringSubmatch(s)
				if len(match) != 4 {
					panic("invalid input")
				}
				log.Printf("  operation: %s %s %s", match[1], match[2], match[3])

				oper2 := operand{old: match[3] == "old"}
				if !oper2.old {
					oper2.num, err = strconv.Atoi(match[3])
					if err != nil {
						panic("invalid int")
					}
				}
				m.operation = op{right: oper2, operator: match[2][0]}
			case test:
				match := testRx.FindStringSubmatch(s)
				if len(match) != 2 {
					panic("invalid input")
				}
				m.testDivisor, err = strconv.Atoi(match[1])
				if err != nil {
					panic("invalid int")
				}
				log.Printf("  test divisor: %d", m.testDivisor)
			case trueCondition, falseCondition:
				match := conditionRx.FindStringSubmatch(s)
				if len(match) != 3 {
					panic("invalid input")
				}
				dst := &m.trueConditionDest
				if i == falseCondition {
					dst = &m.falseConditionDest
				}
				*dst, err = strconv.Atoi(match[2])
				if err != nil {
					panic("invalid int")
				}

				log.Printf("  %s condition dest: %d", match[1], *dst)
			default:
				panic("bad state")
			}
		}

		monkeys = append(monkeys, m)
		log.Print("")
	}

	for round := 1; round <= 20; round++ {
		for _, m := range monkeys {
			maybeLog("monkey %d:", m.index)
			items := m.items
			for !items.Empty() {
				item, err := items.Pop()
				if err != nil {
					panic("bad queue")
				}
				m.inspectCount++

				maybeLog("  Monkey inspects an item with a worry level of %d.", item)

				right := item
				if !m.operation.right.old {
					right = m.operation.right.num
				}
				switch m.operation.operator {
				case '+':
					item += right
					maybeLog("    Worry level increases by %d to %d.", right, item)
				case '*':
					item *= right
					maybeLog("    Worry level is multiplied by %d to %d.", right, item)
				default:
					panic("bad operator")
				}

				item /= 3
				maybeLog("    Monkey gets bored with item. Worry level is divided by 3 to %d.", item)

				notStr := ""
				nextMonkey := m.trueConditionDest
				if item%m.testDivisor != 0 {
					notStr = "not "
					nextMonkey = m.falseConditionDest
				}
				maybeLog("    Current worry level is %sdivisible by %d.", notStr, m.testDivisor)

				monkeys[nextMonkey].items.Push(item)
				maybeLog("    Item with worry level %d is thrown to monkey %d.", item, nextMonkey)
			}
		}
		log.Printf("")

		log.Printf("After round %d, the monkeys are holding items with these worry levels:", round)
		for _, m := range monkeys {
			log.Printf("Monkey %d: %s", m.index, m.items.Join(", "))
		}
		log.Printf("")
	}

	sort.Slice(monkeys, func(i, j int) bool {
		// sort desc
		return monkeys[i].inspectCount > monkeys[j].inspectCount
	})
	topMonkeys := []int{monkeys[0].index, monkeys[1].index}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].index < monkeys[j].index
	})
	for _, m := range monkeys {
		boldStart := ""
		boldEnd := ""
		if slices.Contains(topMonkeys, m.index) {
			boldStart = "\033[1m"
			boldEnd = "\033[m"
		}
		log.Printf("%sMonkey %d inspected items %d times.%s", boldStart, m.index, m.inspectCount, boldEnd)
	}
	log.Printf("")

	log.Printf("monkey business: %d", monkeys[topMonkeys[0]].inspectCount*monkeys[topMonkeys[1]].inspectCount)
}

func maybeLog(msg string, args ...interface{}) {
	if logVerbose {
		log.Printf(msg, args...)
	}
}
