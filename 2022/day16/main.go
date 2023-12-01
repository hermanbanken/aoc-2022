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
var dist map[lib.Vertex]map[lib.Vertex]int

func main() {
	re := regexp.MustCompile(`Valve (.{2}) has flow rate=(\d+); tunnels? leads? to valves? (.*)`)
	var allValves int64 = 0
	var idxAA = 0
	var g = lib.NewGraph()
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
			g.Add(v.idx, valves[dest].idx, 1)
		}
		valves[k] = v
	}
	dist, _ = lib.FloydWarshall(g)

	fmt.Printf("%b\n", allValves)
	s := run(quest{minutesRemaining: 30, pos: idxAA, remainingValves: allValves})
	fmt.Println(s.flow, s.Path(), len(cache), hits)

	cache = map[quest]solution{}
	s = run(quest{minutesRemaining: 26, pos: idxAA, remainingValves: allValves})
	fmt.Println("Done solving on 26", len(cache))

	var pairs [][2]quest
	for k1, v1 := range cache {
		for k2, v2 := range cache {
			if k1 == k2 || v2.flow < s.flow/4 {
				continue
			}
			// fmt.Println(v1, v2)
			offsetA := 0
			offsetB := 0
			if v1.path[0] == idxAA {
				offsetA = 1
			}
			if v2.path[0] == idxAA {
				offsetB = 1
			}
			if len(intersect(v1.path[offsetA:], v2.path[offsetB:])) == 0 {
				// fmt.Println(v1.path, v2.path, intersect(v1.path, v2.path))
				pairs = append(pairs, [2]quest{k1, k2})
			}
		}
	}
	sort.Slice(pairs, func(i, j int) bool {
		a := run(quest{minutesRemaining: 26, pos: idxAA, remainingValves: cache[pairs[i][0]].Available(idxAA)})
		b := run(quest{minutesRemaining: 26, pos: idxAA, remainingValves: cache[pairs[i][1]].Available(idxAA)})
		c := run(quest{minutesRemaining: 26, pos: idxAA, remainingValves: cache[pairs[j][0]].Available(idxAA)})
		d := run(quest{minutesRemaining: 26, pos: idxAA, remainingValves: cache[pairs[j][1]].Available(idxAA)})
		return a.flow+b.flow > c.flow+d.flow
	})
	a := run(quest{minutesRemaining: 26, pos: idxAA, remainingValves: cache[pairs[0][0]].Available(idxAA)})
	b := run(quest{minutesRemaining: 26, pos: idxAA, remainingValves: cache[pairs[0][1]].Available(idxAA)})
	fmt.Println(a.flow + b.flow)
}

func intersect(as, bs []int) (both []int) {
	for _, a := range as {
		for _, b := range bs {
			if a == b {
				both = append(both, a)
			}
		}
	}
	return
}

type quest struct {
	minutesRemaining int
	pos              int
	remainingValves  int64
}

type solution struct {
	flow int
	path []int
}

func (s solution) Available(skip int) (out int64) {
	for p := range s.path {
		if p == skip {
			continue
		}
		out |= 1 << p
	}
	return
}

func (s solution) Path() (out string) {
	for _, p := range s.path {
		if out == "" {
			out = valvesIdx[p]
		} else {
			out = out + " -> " + valvesIdx[p]
		}
	}
	return
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
	v := valves[valvesIdx[q.pos]]
	solutions := []solution{}

	// Current node when opened adds [extra] flow
	var extra int
	var remainingValves int64 = q.remainingValves
	var opened = 0
	if hasValve(remainingValves, q.pos) {
		extra = v.flowrate * (q.minutesRemaining - 1)
		solutions = append(solutions, solution{flow: extra, path: []int{q.pos}})
		remainingValves = removeValve(remainingValves, q.pos)
		opened = 1
	} else {
		solutions = append(solutions, solution{flow: 0, path: []int{}})
	}

	if q.minutesRemaining <= 1 {
		return solutions[0]
	}
	if c, hasCache := cache[q]; hasCache {
		hits++
		if hits%(2<<20) == 0 {
			fmt.Println(hits, len(cache))
		}
		return c
	}

	for i := 0; i < 64; i++ {
		if hasValve(remainingValves, i) {
			dist := dist[lib.Vertex(q.pos)][lib.Vertex(i)]
			s := run(quest{
				minutesRemaining: q.minutesRemaining - dist - opened,
				pos:              i,
				remainingValves:  remainingValves,
			})
			s.flow += extra
			s.path = append([]int{q.pos}, s.path...)
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
