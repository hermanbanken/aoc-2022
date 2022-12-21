package main

import (
	"aoc/lib"
	"log"
	"math/big"
	"strings"
)

type Monkey struct {
	op    string
	elemA string
	elemB string
	sum   *big.Int
}

func (m Monkey) Sum() *big.Int {
	if m.sum != nil {
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

func (m Monkey) MakeEqualTo(nr *big.Int) (human *big.Int) {
	if m.elemA == "" {
		log.Println("Human", m, nr)
		return nr
	}

	log.Println("MakeEqualTo", nr, "(", m.elemA, m.op, m.elemB, ")")
	if m.A().HasHuman(m.elemA) {
		fixed := m.B().Sum()
		dynamic := inverseOp(m.op, nr, nil, fixed)
		if op(m.op, dynamic, fixed).Cmp(nr) != 0 {
			log.Fatal("a math doesnt work out ", op(m.op, dynamic, fixed), "!=", nr, " || ", nr, "==", "[", dynamic, "]", m.op, fixed)
		}
		human = m.A().MakeEqualTo(dynamic)
	} else {
		fixed := m.A().Sum()
		dynamic := inverseOp(m.op, nr, fixed, nil)
		if op(m.op, fixed, dynamic).Cmp(nr) != 0 {
			log.Fatal("b math doesnt work out ", op(m.op, fixed, dynamic), "!=", nr, " || ", nr, "==", fixed, m.op, "[", dynamic, "]")
		}
		human = m.B().MakeEqualTo(dynamic)
	}
	return human
}

func inverseOp(name string, result *big.Int, componentA, componentB *big.Int) (out *big.Int) {
	out = big.NewInt(0)
	component := lib.Ternary(componentA == nil, componentB, componentA)
	if componentB == nil {
		switch name {
		case "*":
			return out.Div(result, component)
		case "+":
			return out.Sub(result, component)
		case "/":
			return out.Mul(result, component)
		case "-":
			return out.Sub(component, result)
		default:
			panic("unknown op: " + name)
		}
	} else {
		switch name {
		case "*":
			return out.Div(result, component)
		case "+":
			return out.Sub(result, component)
		case "/":
			return out.Mul(result, component)
		case "-":
			return out.Add(result, component)
		default:
			panic("unknown op: " + name)
		}
	}
}

func op(name string, a, b *big.Int) (out *big.Int) {
	out = big.NewInt(0)
	switch name {
	case "*":
		return out.Mul(a, b)
	case "+":
		return out.Add(a, b)
	case "/":
		return out.Div(a, b)
	case "-":
		return out.Sub(a, b)
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
			monkeys[line[0:4]] = Monkey{sum: big.NewInt(lib.Int64(line[6:]))}
		}
	})

	log.Println("partA", monkeys["root"].Sum())

	var humn = big.NewInt(0)
	if monkeys["root"].A().HasHuman(monkeys["root"].elemA) {
		humn = monkeys["root"].A().MakeEqualTo(monkeys["root"].B().Sum())
	} else {
		humn = monkeys["root"].B().MakeEqualTo(monkeys["root"].A().Sum())
	}
	// not 8112670909931
	// answer is 3423279932937
	log.Println("partB", humn)
}
