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

type Item int

type Monkey struct {
	items           []Item
	Op              func(Item) Item
	TestDivisibleBy Item
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
		m.items = lib.Map(strings.Split(strings.Split(scanner.Text(), ": ")[1], ", "), func(v string) Item {
			d, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			return Item(d)
		})
		scanner.Scan()
		m.Op = parseOp(strings.TrimPrefix(scanner.Text(), "Operation: "))
		scanner.Scan()
		div, _ := strconv.Atoi(lib.Last(strings.Split(scanner.Text(), " ")))
		m.TestDivisibleBy = Item(div)

		scanner.Scan()
		m.OnTrueMonkey, _ = strconv.Atoi(lib.Last(strings.Split(scanner.Text(), " ")))
		scanner.Scan()
		m.OnFalseMonkey, _ = strconv.Atoi(lib.Last(strings.Split(scanner.Text(), " ")))
		scanner.Scan()

		monkeys = append(monkeys, &m)
		fmt.Println(m)
	}

	var product int64 = 1
	for _, m := range monkeys {
		product *= int64(m.TestDivisibleBy)
	}

	for i := 1; i <= 10000; i++ {
		fmt.Println("round", i)
		for _, m := range monkeys {
			m.Round(monkeys, Item(product), false)
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

func (m *Monkey) Round(others []*Monkey, product Item, divideByThree bool) {
	if len(m.items) == 0 {
		return
	}

	for _, item := range m.items {
		m.business += 1
		var newLevel Item
		if divideByThree {
			newLevel = m.Op(item) / Item(3)
		} else {
			newLevel = m.Op(item) % product
		}
		if newLevel.IsDivisibleBy(m.TestDivisibleBy) {
			others[m.OnTrueMonkey].items = append(others[m.OnTrueMonkey].items, newLevel)
		} else {
			others[m.OnFalseMonkey].items = append(others[m.OnFalseMonkey].items, newLevel)
		}
	}
	m.items = nil
}

func (item Item) IsDivisibleBy(div Item) bool {
	return item%div == 0
}

// lazy math; no attempt to parse needed
func doOp(op string, old Item) (new Item) {
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
	return func(item Item) Item {
		item = doOp(op, item)
		return item
	}
}
