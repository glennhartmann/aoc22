package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/glennhartmann/aoclib/common"
)

type monkey struct {
	name       string
	op         mathOp
	number     int64
	dependents []*monkey
}

type mathOp struct {
	left, right operand
	operator    operator
}

var mathNoop = mathOp{} // zero value of operator ('\0') is not a valid mathOp

type operand struct {
	name   string
	number int64
}

func (op operand) String() string {
	if op.name == "" {
		return fmt.Sprintf("%d", op.number)
	}
	return op.name
}

type operator byte

const (
	add      operator = '+'
	multiply          = '*'
	divide            = '/'
	subtract          = '-'
)

var rx = regexp.MustCompile(`(\w+): ((\d+)|((\w+) (.) (\w+)))`)

func main() {
	monkeys := make(map[string]*monkey, 50)
	deps := make(map[string][]*monkey, 50)
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
		if len(match) != 8 {
			panic("invalid input or regex")
		}

		m := &monkey{name: match[1]}
		log.Printf("%s:", m.name)
		num := match[3]
		if num == "" {
			m.op = mathOp{
				left:     operand{name: match[5]},
				right:    operand{name: match[7]},
				operator: operator(match[6][0]),
			}
			log.Printf("  mathOp: %s %c %s", m.op.left.String(), byte(m.op.operator), m.op.right.String())
			handleDeps(monkeys, deps, m)
		} else {
			m.number, err = strconv.ParseInt(num, 10, 64)
			if err != nil {
				panic("bad int64")
			}
			log.Printf("  const: %d", m.number)
		}

		log.Printf("  adding dependent monkeys from global list to my dependents list: %s", depStr(deps[m.name]))
		m.dependents = append(m.dependents, deps[m.name]...)
		log.Printf("  new dependents list: %s. Maybe notifying them...", depStr(m.dependents))
		m.maybeNotifyDependents("    ")

		monkeys[m.name] = m
	}
	log.Printf("root: %d", monkeys["root"].number)
}

func handleDeps(monkeys map[string]*monkey, deps map[string][]*monkey, m *monkey) {
	resolve := func(lr *monkey, operand *operand) {
		if lr.op == mathNoop {
			operand.name = ""
			operand.number = lr.number
			log.Printf("    resolved as number: %d", operand.number)
		} else {
			log.Printf("    adding self (%q) to %q's list of dependents", m.name, lr.name)
			lr.dependents = append(lr.dependents, m)
		}
	}

	if left := monkeys[m.op.left.name]; left != nil {
		log.Printf("  left operand (%q) already known", left.name)
		resolve(left, &m.op.left)
	} else {
		log.Printf("  adding left operand to global dependents map {%q: [..., %q]} (ie, %q depends on %q)", m.op.left.name, m.name, m.name, m.op.left.name)
		deps[m.op.left.name] = append(deps[m.op.left.name], m)
		log.Printf("  %q's new list of depentents: %s", m.op.left.name, depStr(deps[m.op.left.name]))
	}
	if right := monkeys[m.op.right.name]; right != nil {
		log.Printf("  right operand (%q) already known", right.name)
		resolve(right, &m.op.right)
	} else {
		log.Printf("  adding right operand to global dependents map {%q: [..., %q]} (ie, %q depends on %q)", m.op.right.name, m.name, m.name, m.op.right.name)
		deps[m.op.right.name] = append(deps[m.op.right.name], m)
		log.Printf("  %q's new list of depentents: %s", m.op.right.name, depStr(deps[m.op.left.name]))
	}

	if m.op.left.name == "" && m.op.right.name == "" {
		log.Printf("  both operands resolved to numbers - computing result and notifying %d dependents...", len(m.dependents))
		m.compute("  ")
		m.maybeNotifyDependents("    ")
	}
}

func depStr(deps []*monkey) string {
	return fmt.Sprintf("[%s]", common.Fjoin(deps, ", ", func(dep *monkey) string { return fmt.Sprintf("%q", dep.name) }))
}

func (m *monkey) maybeNotifyDependents(spaces string) {
	if m.op != mathNoop {
		log.Printf("%snot a const - not notifying dependents", spaces)
		return
	}
	log.Printf("%snotifying %q's %d dependents", spaces, m.name, len(m.dependents))
	for _, dep := range m.dependents {
		if dep.op.left.name == m.name {
			dep.op.left.name = ""
			dep.op.left.number = m.number
			log.Printf("%s  %q.op.left = %d", spaces, dep.name, m.number)
			if dep.op.right.name == "" {
				dep.compute(spaces + "  ")
				dep.maybeNotifyDependents(spaces + "  ")
				continue
			}
		}
		if dep.op.right.name == m.name {
			dep.op.right.name = ""
			dep.op.right.number = m.number
			log.Printf("%s  %q.op.right = %d", spaces, dep.name, m.number)
			if dep.op.left.name == "" {
				dep.compute(spaces + "  ")
				dep.maybeNotifyDependents(spaces + "  ")
			}
		}
	}
	m.dependents = nil
}

func (m *monkey) compute(spaces string) {
	if m.op.left.name != "" || m.op.right.name != "" {
		panic("can't compute - not const")
	}

	switch m.op.operator {
	case add:
		m.number = m.op.left.number + m.op.right.number
	case multiply:
		m.number = m.op.left.number * m.op.right.number
	case divide:
		m.number = m.op.left.number / m.op.right.number
	case subtract:
		m.number = m.op.left.number - m.op.right.number
	default:
		panic("bad operator")
	}
	log.Printf("%s%q computed as %d %c %d = %d", spaces, m.name, m.op.left.number, byte(m.op.operator), m.op.right.number, m.number)
	m.op = mathNoop
}
