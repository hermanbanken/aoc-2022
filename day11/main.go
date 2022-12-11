package main

import (
	"aoc/lib"
	"bufio"
	"fmt"
	"log"
	"math/big"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items           []*big.Int
	Op              func(*big.Int) *big.Int
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
		m.items = doMap(strings.Split(strings.Split(scanner.Text(), ": ")[1], ", "), func(v string) *big.Int {
			d, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			return big.NewInt(int64(d))
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
		for _, m := range monkeys {
			m.Round(monkeys)
		}
		fmt.Printf("== After round %d ==\n", i)
		for i, m := range monkeys {
			fmt.Printf("Monkey %d inspected items %d times.\n", i, m.business)
			// fmt.Printf("Monkey %d: %s\n", i, strings.Join(doMap(m.items, func(v int) string { return strconv.Itoa(v) }), ", "))
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
		newLevel := m.Op(item) // / 3
		if big.NewInt(0).Mod(newLevel, big.NewInt(int64(m.TestDivisibleBy))).Int64() == 0 {
			others[m.OnTrueMonkey].items = append(others[m.OnTrueMonkey].items, newLevel)
		} else {
			others[m.OnFalseMonkey].items = append(others[m.OnFalseMonkey].items, newLevel)
		}
	}
	m.items = nil
}

func parseOp(op string) func(*big.Int) *big.Int {
	op = strings.TrimSpace(op)
	op = strings.TrimPrefix(op, "Operation: ")
	return func(old *big.Int) (new *big.Int) {
		switch op {
		case "new = old * 7":
			return big.NewInt(0).Mul(old, big.NewInt(7))
		case "new = old + 8":
			return big.NewInt(0).Add(old, big.NewInt(8))
		case "new = old + 7":
			return big.NewInt(0).Add(old, big.NewInt(7))
		case "new = old + 2":
			return big.NewInt(0).Add(old, big.NewInt(2))
		case "new = old + 1":
			return big.NewInt(0).Add(old, big.NewInt(1))
		case "new = old + 4":
			return big.NewInt(0).Add(old, big.NewInt(4))
		case "new = old + 6":
			return big.NewInt(0).Add(old, big.NewInt(6))
		case "new = old + 3":
			return big.NewInt(0).Add(old, big.NewInt(3))
		case "new = old * 13":
			return big.NewInt(0).Mul(old, big.NewInt(13))
		case "new = old * old":
			return big.NewInt(0).Mul(old, old)
		case "new = old * 19":
			return big.NewInt(0).Mul(old, big.NewInt(19))
		default:
			panic("unknown operation: " + op)
		}
		return
	}

}

func doMap[T any, R any](ts []T, fn func(T) R) (out []R) {
	out = make([]R, len(ts))
	for i, t := range ts {
		out[i] = fn(t)
	}
	return
}
