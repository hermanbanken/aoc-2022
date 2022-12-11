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

type Monkey struct {
	items           []int
	OpOp            byte
	OpOther         int
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
		op := strings.Split(strings.TrimPrefix(scanner.Text(), "  Operation: new = old "), " ")
		m.OpOp = op[0][0]
		m.OpOther, _ = strconv.Atoi(op[1])
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

func (m Monkey) Calc(old int) int {
	d := lib.Ternary(m.OpOther == 0 /* "old" */, old, m.OpOther)
	return lib.Ternary(m.OpOp == '+', old+d, old*d)
}

func (m *Monkey) Round(others []*Monkey, product int, divideByThree bool) {
	for _, item := range m.items {
		m.business += 1
		var newLevel = m.Calc(item)
		if divideByThree {
			newLevel /= 3
		} else {
			newLevel %= product
		}
		target := lib.Ternary(newLevel%m.TestDivisibleBy == 0, m.OnTrueMonkey, m.OnFalseMonkey)
		others[target].items = append(others[m.OnTrueMonkey].items, newLevel)
	}
	m.items = nil
}
