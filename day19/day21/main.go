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

func (m Monkey) A() Monkey { return monkeys[m.elemA] }
func (m Monkey) B() Monkey { return monkeys[m.elemB] }

func (m Monkey) Sum() int64 {
	if m.sum != 0 {
		return m.sum
	}
	return op(m.op, m.A().Sum(), m.B().Sum())
}

func (m Monkey) HasHuman(name string) bool {
	if name == "humn" || m.elemA == "humn" || m.elemB == "humn" {
		return true
	}
	if m.elemA == "" {
		return false
	}
	return m.A().HasHuman(m.elemA) || m.B().HasHuman(m.elemB)
}

func (m Monkey) MakeEqualTo(nr int64) (human int64) {
	if m.elemA == "" {
		log.Println("Human", m, nr)
		return nr
	}

	log.Printf("MakeEqualTo: %d == %s %s %s\n", nr, m.elemA, m.op, m.elemB)
	var a, b int64
	if m.A().HasHuman(m.elemA) {
		b = m.B().Sum()
		a = inverseOp(m.op, nr, b, true)
	} else {
		a = m.A().Sum()
		b = inverseOp(m.op, nr, a, false)
	}
	if op(m.op, a, b) != nr {
		log.Fatalf(lib.Ternary(m.A().HasHuman(m.elemA),
			"math doesnt work out %d != %d (math: %d == [%d] %s %d)",
			"math doesnt work out %d != %d (math: %d == %d %s [%d])"),
			op(m.op, a, b), nr, nr, a, m.op, b)
	}

	if m.A().HasHuman(m.elemA) {
		return m.A().MakeEqualTo(a)
	} else {
		return m.B().MakeEqualTo(b)
	}
}

func inverseOp(name string, result int64, component int64, calcLeftComponent bool) (out int64) {
	switch name {
	case "*":
		return result / component
	case "+":
		return result - component
	case "/":
		if calcLeftComponent {
			return result * component
		}
		return component / result
	case "-":
		if !calcLeftComponent {
			return component - result
		}
		return result + component
	default:
		panic("unknown op: " + name)
	}
}

func op(name string, a, b int64) (out int64) {
	switch name {
	case "*":
		return (a * b)
	case "+":
		return (a + b)
	case "/":
		return (a / b)
	case "-":
		return (a - b)
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

	if monkeys["root"].A().HasHuman(monkeys["root"].elemA) {
		log.Println("partB", monkeys["root"].A().MakeEqualTo(monkeys["root"].B().Sum()))
	} else {
		log.Println("partB", monkeys["root"].B().MakeEqualTo(monkeys["root"].A().Sum()))
	}
}
