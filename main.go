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
	allValves := ""
	lib.EachLine(func(line string) {
		m := re.FindStringSubmatch(line)
		if m == nil {
			fmt.Println("no match", line)
		}
		valve := m[1]
		fr := lib.Int(m[2])
		dest := strings.Split(m[3], ", ")
		valves[valve] = Valve{len(valves), valve, fr, dest}
		if fr > 0 {
			allValves += valve
		}
	})

	fmt.Println(allValves)
	s := run(quest{30, "AA", allValves})
	fmt.Println(s, len(cache), hits)
}

type quest struct {
	minutesRemaining int
	pos              string
	remainingValves  string
}

func flowrate(vs string) (sum int) {
	for i := 0; i < len(vs)/2; i++ {
		sum += valves[vs[i*2:i*2+2]].flowrate
	}
	return sum
}

type solution struct {
	valves string
	flow   int
}

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

func removeValve(vs string, v string) string {
	src := []byte(vs)
	dst := make([]byte, len(vs)-2)

	for i := 0; i < len(src)/2; i += 1 {
		if vs[i*2:i*2+2] == v {
			copy(dst, src[0:i*2])
			copy(dst[i*2:], src[i*2+2:])
			return string(dst)
		}
	}
	panic("cant remove valve")
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
var cache map[quest]solution = map[quest]solution{}

func run(q quest) (sol solution) {
	if q.minutesRemaining <= 1 {
		return solution{flow: 0}
	}
	if c, hasCache := cache[q]; hasCache {
		return c
	}

	v := valves[q.pos]
	solutions := []solution{}

	for _, shouldOpen := range []bool{true, false} {
		var extra = 0
		var remaining = q.remainingValves
		if shouldOpen {
			if !hasValve(q.remainingValves, q.pos) {
				continue
			}
			extra = lib.Ternary(shouldOpen, v.flowrate*(q.minutesRemaining-1), 0)
			remaining = lib.Ternary(shouldOpen, removeValve(q.remainingValves, q.pos), q.remainingValves)
		}

		solutions = append(solutions, solution{flow: extra})
		for _, dest := range v.dest {
			// fmt.Println("moving", dest)
			s := run(quest{
				minutesRemaining: q.minutesRemaining - 1 - lib.Ternary(shouldOpen, 1, 0),
				pos:              dest,
				remainingValves:  remaining,
			})
			s.flow += extra
			solutions = append(solutions, s)
		}
	}

	sort.Slice(solutions, func(i, j int) bool {
		return solutions[i].flow > solutions[j].flow
	})

	if len(solutions) > 0 {
		sol = solutions[0]
	}
	// best: DDBBJJHHEECC
	// if sol.flow > 0 {
	// 	// fmt.Println(q.minutesRemaining, q.pos, len(q.remainingValves)/2, sol.flow)
	// }
	cache[q] = sol
	return sol
}
