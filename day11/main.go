package main

import (
	"aoc/lib"
	"bufio"
	"fmt"
	"log"
	"sort"
	"strings"
)

type Monkey struct {
	items           []int
	Op              func(int) int
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
		m.items = lib.Map(strings.Split(strings.Split(scanner.Text(), ": ")[1], ", "), func(v string) int {
			return lib.Int(v)
		})
		scanner.Scan()
		op := strings.TrimPrefix(scanner.Text(), "  Operation: ")
		m.Op = func(i int) int { return doOp(op, i) }
		scanner.Scan()
		m.TestDivisibleBy = lib.Int(lib.Last(strings.Split(scanner.Text(), " ")))
		scanner.Scan()
		m.OnTrueMonkey = lib.Int(lib.Last(strings.Split(scanner.Text(), " ")))
		scanner.Scan()
		m.OnFalseMonkey = lib.Int(lib.Last(strings.Split(scanner.Text(), " ")))
		scanner.Scan()

		monkeys = append(monkeys, &m)
		fmt.Println(m)
	}

	var product int = 1
	for _, m := range monkeys {
		product *= int(m.TestDivisibleBy)
	}

	for i := 1; i <= 10000; i++ {
		fmt.Println("round", i)
		for _, m := range monkeys {
			m.Round(monkeys, product, false)
		}
		for i, m := range monkeys {
			fmt.Printf("Monkey %d inspected items %d times.\n", i, m.business)
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].business > monkeys[j].business
	})
	fmt.Println(monkeys[0].business * monkeys[1].business)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (m *Monkey) Round(others []*Monkey, product int, divideByThree bool) {
	for _, item := range m.items {
		m.business += 1
		var newLevel = m.Op(item)
		if divideByThree {
			newLevel /= 3
		} else {
			newLevel %= product
		}
		if newLevel%m.TestDivisibleBy == 0 {
			others[m.OnTrueMonkey].items = append(others[m.OnTrueMonkey].items, newLevel)
		} else {
			others[m.OnFalseMonkey].items = append(others[m.OnFalseMonkey].items, newLevel)
		}
	}
	m.items = nil
}

// lazy math; no attempt to parse needed
func doOp(op string, old int) (new int) {
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
