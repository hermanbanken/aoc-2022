package main

import (
	"aoc/lib"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type Valve struct {
	idx      int
	id       string
	flowrate int
	dest     []string
}

var valves map[string]Valve = map[string]Valve{}

func main() {
	re := regexp.MustCompile(`Valve (.{2}) has flow rate=(\d+); tunnels? leads? to valves? (.*)`)
	lib.EachLine(func(line string) {
		m := re.FindStringSubmatch(line)
		if m == nil {
			fmt.Println("no match", line)
		}
		valve := m[1]
		fr := lib.Int(m[2])
		dest := strings.Split(m[3], ", ")
		valves[valve] = Valve{len(valves), valve, fr, dest}
	})

	s := run(sub{30, "AA", ""})
	fmt.Println(s, len(cache), hits)
}

type sub struct {
	minutesRemaining int
	pos              string
	valves           string
}

type solution struct {
	valves string
	flow   int
}

var cache map[sub]solution = map[sub]solution{}

func addValve(vs string, v string) string {
	return vs + v
	// items := []string{}
	// for i := 0; i < len(vs)-1; i += 2 {
	// 	items = append(items, vs[i:i+2])
	// }
	// items = append(items, v)
	// sort.Strings(items)
	// return strings.Join(items, "")
}

func hasValve(vs string, v string) bool {
	for i := 0; i < len(vs)-1; i += 2 {
		if vs[i:i+2] == v {
			return true
		}
	}
	return false
}

var hits = 0

func run(q sub) (sol solution) {
	if q.minutesRemaining <= 1 {
		return solution{flow: 0, valves: q.valves}
	}
	if max, hasCache := cache[q]; hasCache {
		hits += 1
		return max
	}

	v := valves[q.pos]
	solutions := []solution{}

	for _, shouldOpen := range []bool{true, false} {
		if hasValve(q.valves, q.pos) {
			continue
		}
		extra := lib.Ternary(shouldOpen, v.flowrate*(q.minutesRemaining-1), 0)
		valves := lib.Ternary(shouldOpen, addValve(q.valves, q.pos), q.valves)
		solutions = append(solutions, solution{valves, extra})
		for _, dest := range v.dest {
			s := run(sub{minutesRemaining: q.minutesRemaining - 1 - lib.Ternary(shouldOpen, 1, 0), pos: dest, valves: valves})
			s.flow += extra
			solutions = append(solutions, s)
		}
	}

	sort.Slice(solutions, func(i, j int) bool {
		return solutions[i].flow > solutions[j].flow
	})

	cache[q] = solution{q.valves, 0}
	if len(solutions) > 0 {
		cache[q] = solutions[0]
	}
	// best: DDBBJJHHEECC
	// if cache[q].flow > 0 {
	// 	fmt.Println(q.minutesRemaining, q.pos, q.valves, cache[q].valves, cache[q].flow)
	// }
	return cache[q]
}
