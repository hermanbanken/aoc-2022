// / Suboptimal solution that runs up to 30 minutes and uses more than 20GB of RAM to memoize.
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
	destIdx  []int
}

var valves map[string]Valve = map[string]Valve{}
var valvesIdx map[int]string = map[int]string{}

func main() {
	re := regexp.MustCompile(`Valve (.{2}) has flow rate=(\d+); tunnels? leads? to valves? (.*)`)
	var allValves int64 = 0
	var idxAA = 0
	lib.EachLine(func(line string) {
		m := re.FindStringSubmatch(line)
		if m == nil {
			fmt.Println("no match", line)
		}
		valve := m[1]
		fr := lib.Int(m[2])
		dest := strings.Split(m[3], ", ")
		valvesIdx[len(valves)] = valve
		if fr > 0 {
			allValves |= 1 << len(valves)
		}
		if valve == "AA" {
			idxAA = len(valves)
		}
		valves[valve] = Valve{len(valves), valve, fr, dest, nil}
	})
	for k, v := range valves {
		for _, dest := range v.dest {
			v.destIdx = append(v.destIdx, valves[dest].idx)
		}
		valves[k] = v
	}
	fmt.Println(valves)

	fmt.Printf("%b\n", allValves)
	s := run(quest{minutesRemaining: 26, pos: [2]int{idxAA, idxAA}, remainingValves: allValves})
	fmt.Println(s, len(cache), hits)
}

type quest struct {
	minutesRemaining int
	pos              [2]int
	blocked          [2]bool
	remainingValves  int64
}

type solution struct {
	flow int
}

func removeValve(vs int64, v int) int64 {
	if !hasValve(vs, v) {
		panic("cant remove valve")
	}
	var mask int64 = ^(1 << v)
	vs &= mask
	if hasValve(vs, v) {
		panic("remove valve failed")
	}
	return vs
}

func hasValve(vs int64, v int) bool {
	return vs&(1<<v) > 0
}

var hits = 0
var cache map[quest]solution = map[quest]solution{}

func run(q quest) (sol solution) {
	if q.minutesRemaining <= 1 {
		return solution{flow: 0}
	}
	if c, hasCache := cache[q]; hasCache {
		hits++
		if hits%(2<<20) == 0 {
			fmt.Println(hits, len(cache))
		}
		return c
	}

	p1 := q.pos[0]
	p2 := q.pos[1]
	v1 := valves[valvesIdx[p1]]
	v2 := valves[valvesIdx[p2]]
	solutions := []solution{}

	for _, shouldOpen := range [][]bool{{true, true}, {true, false}, {false, true}, {false, false}} {
		var extra = 0
		var remaining = q.remainingValves
		opens := []bool{!q.blocked[0] && shouldOpen[0], !q.blocked[1] && shouldOpen[1]}
		if opens[0] {
			if !hasValve(remaining, p1) {
				continue
			}
			extra += v1.flowrate * (q.minutesRemaining - 1)
			remaining = removeValve(remaining, p1)
		}
		if opens[1] {
			if !hasValve(remaining, p2) {
				continue
			}
			extra += v2.flowrate * (q.minutesRemaining - 1)
			remaining = removeValve(remaining, p2)
		}

		solutions = append(solutions, solution{flow: extra})
		var blocked [2]bool
		possiblePos := [][2]int{}
		for _, dest1 := range lib.Ternary(opens[0], []int{p1}, v1.destIdx) {
			for _, dest2 := range lib.Ternary(opens[1], []int{p2}, v2.destIdx) {
				if dest1 < dest2 {
					possiblePos = append(possiblePos, [2]int{dest1, dest2})
					blocked = [2]bool{opens[0], opens[1]}
				} else {
					possiblePos = append(possiblePos, [2]int{dest2, dest1})
					blocked = [2]bool{opens[1], opens[0]}
				}
			}
		}
		// fmt.Println("possiblePos", possiblePos)
		possiblePos = lib.Unique(possiblePos, func(a, b [2]int) bool { return lib.Ternary(a[0] == b[0], a[1]-b[1], a[0]-b[0]) > 0 })
		// fmt.Println("possiblePos", possiblePos)

		for _, dest := range possiblePos {
			s := run(quest{
				minutesRemaining: q.minutesRemaining - 1,
				pos:              dest,
				blocked:          blocked,
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
