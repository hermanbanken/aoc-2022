package main

import (
	"aoc/lib"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

func main() {
	var lines []string

	// Part 1
	var sumIndexes = 0
	lib.EachLine(func(line string) {
		if strings.TrimSpace(line) == "" {
			return
		}
		lines = append(lines, line)
		if len(lines) > 0 && len(lines)%2 == 0 {
			if compare(parse(lines[len(lines)-2]), parse(lines[len(lines)-1])) <= 0 {
				sumIndexes += len(lines) / 2
			}
		}
	})
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
}

func parse(str string) (out interface{}) {
	lib.Must(json.Unmarshal([]byte(str), &out))
	return
}

func compare(a interface{}, b interface{}) int {
	fa, AisFloat := a.(float64)
	fb, BisFloat := b.(float64)
	if AisFloat && BisFloat {
		return compareInt(int(fa), int(fb))
	}
	la := asList(a)
	lb := asList(b)

	for len(lb) > 0 {
		if len(la) == 0 {
			return -1 // a ran out
		}
		// fmt.Println("compare", la[0], lb[0])
		if result := compare(la[0], lb[0]); result != 0 {
			return result
		}
		la = la[1:]
		lb = lb[1:]
	}
	if len(la) == 0 {
		return 0 // equal, continue
	}
	return 1 // a continues, b ran out
}

func compareInt(a, b int) int {
	// fmt.Println("compareInt", a, b)
	return lib.Ternary(a > b, 1, 0) - lib.Ternary(a < b, 1, 0)
}

func asList(elem interface{}) []interface{} {
	switch v := elem.(type) {
	case []interface{}:
		return v
	default:
		return []interface{}{v}
	}
}
