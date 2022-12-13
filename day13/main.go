package main

import (
	"aoc/lib"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"sort"
)

func main() {
	r := lib.Reader()
	defer r.Close()
	scanner := bufio.NewScanner(r)

	var index = 1
	var sumIndexes = 0
	var lines []string
	for scanner.Scan() {
		a := scanner.Text()
		scanner.Scan()
		b := scanner.Text()
		scanner.Scan()
		lines = append(lines, a, b)

		// Part 1
		if compare(parse(a), parse(b)) <= 0 {
			sumIndexes += index
		}
		index++
	}
	fmt.Println(sumIndexes)

	// Part 2
	lines = append(lines, "[[2]]", "[[6]]")
	sort.Slice(lines, func(i, j int) bool {
		return compare(parse(lines[i]), parse(lines[j])) <= 0
	})
	fmt.Println("\nlines:")
	var indexA, indexB int
	for i, l := range lines {
		// fmt.Println(i, ":", l)
		if l == "[[2]]" {
			indexA = i + 1
		}
		if l == "[[6]]" {
			indexB = i + 1
		}
	}
	fmt.Println((indexA) * (indexB))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func parse(str string) (out interface{}) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e, str)
		}
	}()

	err := json.Unmarshal([]byte(str), &out)
	if err != nil {
		panic(err)
	}
	return
}

func compare(a interface{}, b interface{}) int {
	fa, AisFloat := a.(float64)
	fb, BisFloat := b.(float64)
	if AisFloat && BisFloat {
		return compareF(int(fa), int(fb))
	}
	la := asList(a)
	lb := asList(b)

	for len(lb) > 0 {
		if len(la) == 0 {
			return -1
		}
		// fmt.Println("compare", la[0], lb[0])
		if result := compare(la[0], lb[0]); result != 0 {
			return result
		}
		la = la[1:]
		lb = lb[1:]
	}
	if len(la) == 0 {
		return 0
	}
	// fmt.Println("a continues")
	return 1
}

func compareF(a, b int) int {
	// fmt.Println("compareF", a, b)
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

func asList(elem interface{}) []interface{} {
	switch v := elem.(type) {
	case []interface{}:
		return v
	default:
		return []interface{}{v}
	}
}
