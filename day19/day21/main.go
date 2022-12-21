package main

import (
	"aoc/lib"
	"log"
	"strings"
)

type Monkey struct {
	op    string
	elemA string
	elemB string
	sum   int64
}

func (m Monkey) Sum() int64 {
	if m.sum != 0 {
		return m.sum
	}
	a := monkeys[m.elemA].Sum()
	b := monkeys[m.elemB].Sum()
	return op(m.op, a, b)
}

func (m Monkey) A() Monkey {
	return monkeys[m.elemA]
}
func (m Monkey) B() Monkey {
	return monkeys[m.elemB]
}

func (m Monkey) HasHuman(name string) bool {
	if name == "humn" || m.elemA == "humn" || m.elemB == "humn" {
		return true
	}
	if m.elemA == "" {
		return false
	}
	return monkeys[m.elemA].HasHuman(m.elemA) || monkeys[m.elemB].HasHuman(m.elemB)
}

func (m Monkey) MakeEqualTo(nr int64) (human int64) {
	if m.elemA == "" {
		log.Println("Human", m, nr)
		return nr
	}

	log.Println("MakeEqualTo", nr, "(", m.elemA, m.op, m.elemB, ")")
	if m.A().HasHuman(m.elemA) {
		fixed := m.B().Sum()
		dynamic := inverseOp(m.op, nr, fixed)
		if op(m.op, dynamic, fixed) != nr {
			log.Fatal("a math doesnt work out ", op(m.op, dynamic, fixed), "!=", nr, "||", dynamic, m.op, fixed)
		}
		human = m.A().MakeEqualTo(dynamic)
	} else {
		fixed := m.A().Sum()
		dynamic := inverseOp(m.op, nr, fixed)
		if op(m.op, fixed, dynamic) != nr {
			log.Fatal("b math doesnt work out ", op(m.op, fixed, dynamic), "!=", nr, "||", fixed, m.op, dynamic)
		}
		human = m.B().MakeEqualTo(dynamic)
	}
	return human
}

func inverseOp(name string, result int64, component int64) int64 {
	switch name {
	case "*":
		return result / component
	case "+":
		return result - component
	case "/":
		return result * component
	case "-":
		return result + component
	default:
		panic("unknown op: " + name)
	}
}

func op(name string, a, b int64) int64 {
	switch name {
	case "*":
		return a * b
	case "+":
		return a + b
	case "/":
		return a / b
	case "-":
		return a - b
	default:
		panic("unknown op: " + name)
	}
}

var monkeys = map[string]Monkey{}

func main() {
	lib.EachLine(func(line string) {
		if strings.ContainsRune(line[6:], ' ') {
			parts := strings.Split(line[6:], " ")
			monkeys[line[0:4]] = Monkey{op: parts[1], elemA: parts[0], elemB: parts[2]}
		} else {
			if lib.Int(line[6:]) == 0 {
				panic("0 in input")
			}
			monkeys[line[0:4]] = Monkey{sum: lib.Int64(line[6:])}
		}
	})

	log.Println("partA", monkeys["root"].Sum())

	var humn = int64(0)
	if monkeys["root"].A().HasHuman(monkeys["root"].elemA) {
		humn = monkeys["root"].A().MakeEqualTo(monkeys["root"].B().Sum())
	} else {
		humn = monkeys["root"].B().MakeEqualTo(monkeys["root"].A().Sum())
	}
	// not 8112670909931
	log.Println("partB", humn)
}
