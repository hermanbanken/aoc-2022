package main

import (
	"aoc/lib"
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Monkey struct {
	items           []int
	Op              func(int) int
	TestDivisibleBy int
	OnTrueMonkey    int
	OnFalseMonkey   int
}

func main() {
	r := lib.Reader()
	defer r.Close()
	scanner := bufio.NewScanner(r)

	var monkeys []Monkey
	for scanner.Scan() {
		_ = scanner.Text()
		var m Monkey

		scanner.Scan()
		m.items = doMap(strings.Split(strings.Split(scanner.Text(), ": ")[1], ", "), func(v string) int {
			d, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			return d
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

		monkeys = append(monkeys, m)
		fmt.Println(m)
	}

	fmt.Println(monkeys)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func parseOp(op string) func(int) int {
	// TODO
	return func(i int) int { return i }
}

func doMap[T any, R any](ts []T, fn func(T) R) (out []R) {
	out = make([]R, len(ts))
	for i, t := range ts {
		out[i] = fn(t)
	}
	return
}
