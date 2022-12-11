package main

import (
	"aoc/lib"
	"bufio"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

type Item struct {
	base int64
	ops  []*Op
}

type Op struct {
	str      string
	isMult   bool
	isPlus   bool
	isSquare bool
	b        int
}

type Monkey struct {
	items           []*Item
	Op              func(Item) Item
	TestDivisibleBy int
	OnTrueMonkey    int
	OnFalseMonkey   int
	business        int
}

func main() {
	r := lib.Reader()
	defer r.Close()
	scanner := bufio.NewScanner(r)

	var monkeys []*Monkey
	for scanner.Scan() {
		_ = scanner.Text()
		var m Monkey

		scanner.Scan()
		m.items = doMap(strings.Split(strings.Split(scanner.Text(), ": ")[1], ", "), func(v string) *Item {
			d, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			return &Item{base: int64(d)}
		})
		scanner.Scan()
		m.Op = parseOp(strings.TrimPrefix(scanner.Text(), "Operation: "))
		scanner.Scan()
		m.TestDivisibleBy, _ = strconv.Atoi(lib.Last(strings.Split(scanner.Text(), " ")))

		scanner.Scan()
		m.OnTrueMonkey, _ = strconv.Atoi(lib.Last(strings.Split(scanner.Text(), " ")))
		scanner.Scan()
		m.OnFalseMonkey, _ = strconv.Atoi(lib.Last(strings.Split(scanner.Text(), " ")))
		scanner.Scan()

		monkeys = append(monkeys, &m)
		fmt.Println(m)
	}

	for i := 1; i <= 10000; i++ {
		fmt.Println("round", i)
		for _, m := range monkeys {
			m.Round(monkeys)
		}
		for i, m := range monkeys {
			fmt.Printf("Monkey %d inspected items %d times.\n", i, m.business)
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].business > monkeys[j].business
	})
	fmt.Println("part1", monkeys[0].business*monkeys[1].business)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (m *Monkey) Round(others []*Monkey) {
	if len(m.items) == 0 {
		return
	}

	for _, item := range m.items {
		m.business += 1
		newLevel := m.Op(*item)
		if newLevel.IsDivisibleBy(m.TestDivisibleBy) {
			others[m.OnTrueMonkey].items = append(others[m.OnTrueMonkey].items, &newLevel)
		} else {
			others[m.OnFalseMonkey].items = append(others[m.OnFalseMonkey].items, &newLevel)
		}
	}
	m.items = nil
}

func (item Item) IsDivisibleBy(div int) bool {
	b := item.base
	for i := len(item.ops) - 1; i >= 0; i-- {
		if item.ops[i].isMult && item.ops[i].b == div {
			return true
		} else if item.ops[i].isSquare {
			continue
		} else {
			// oh no
			fmt.Println("IsDivisibleBy", b, div, *item)
			panic("dont know what to do now")
		}
	}
	// for _, op := range item.ops {
	// 	b = doOp(op.str, b)
	// }
	fmt.Println("IsDivisibleBy", b, div)
	return b%int64(div) == 0
}

func doOp(op string, old int64) (new int64) {
	switch op {
	case "new = old + 1":
		new = old + 1
	case "new = old + 2":
		new = old + 2
	case "new = old + 3":
		new = old + 3
	case "new = old + 4":
		new = old + 4
	case "new = old + 6":
		new = old + 6
	case "new = old + 7":
		new = old + 7
	case "new = old + 8":
		new = old + 8

	case "new = old * 7":
		new = old * 7
	case "new = old * 13":
		new = old * 13
	case "new = old * 19":
		new = old * 19

	case "new = old * old":
		new = old * old
	default:
		panic("unknown operation: " + op)
	}
	return
}

func parseOp(op string) func(Item) Item {
	op = strings.TrimSpace(op)
	op = strings.TrimPrefix(op, "Operation: ")
	o := Op{str: op}
	o.isMult = strings.Contains(op, "*")
	o.isSquare = strings.Contains(op, "* old")
	o.isPlus = strings.Contains(op, "+")
	o.b, _ = strconv.Atoi(lib.Last(strings.Split(op, " ")))
	return func(item Item) Item {
		item.ops = append(item.ops, &o)
		return item
	}
}

func doMap[T any, R any](ts []T, fn func(T) R) (out []R) {
	out = make([]R, len(ts))
	for i, t := range ts {
		out[i] = fn(t)
	}
	return
}
