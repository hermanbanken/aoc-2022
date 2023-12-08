package main

import (
	"aoc/lib"
	"fmt"
	"strings"
	"time"
)

func main() {
	lines := lib.Lines()
	instructions := lines[0]
	_ = instructions
	lines = lines[2:]

	nodeCount := 0
	nodes := make(map[string]int)
	nodesReverse := make(map[int]string)
	graph := make(map[int][2]int)

	nr := func(str string) int {
		v, ok := nodes[str]
		if !ok {
			nodes[str] = nodeCount
			nodesReverse[nodeCount] = str
			v = nodeCount
			nodeCount++
		}
		return v
	}

	for _, line := range lines {
		parts := strings.FieldsFunc(line, func(r rune) bool { return r == ' ' || r == '=' || r == '(' || r == ',' || r == ')' })
		fmt.Println(parts)
		graph[nr(parts[0])] = [2]int{nr(parts[1]), nr(parts[2])}
	}

	var pos = []int{}
	var endIndexed = map[int]bool{}
	if 1 > 0 {
		for nr, str := range nodesReverse {
			if strings.HasSuffix(str, "A") {
				pos = append(pos, nr)
			} else if strings.HasSuffix(str, "Z") {
				endIndexed[nr] = true
			}
		}
	}

	debug := func() {
		for _, p := range pos {
			fmt.Print(nodesReverse[p], " ")
		}
		fmt.Println()
		time.Sleep(100 * time.Millisecond)
	}
	_ = debug

	i := 0
	var history = make([][]struct {
		T  int
		NR int
	}, len(pos))
	var hasVisited = make([]map[int]bool, len(pos))
	for j := range hasVisited {
		hasVisited[j] = make(map[int]bool)
	}
	var disabled = make([]bool, len(pos))
	disabledCount := 0

	for {
		// part2: detect repeats
		for j, p := range pos {
			if !disabled[j] && strings.HasSuffix(nodesReverse[p], "Z") {
				history[j] = append(history[j], struct {
					T  int
					NR int
				}{i, p})
				if _, ok := hasVisited[j][p]; ok {
					disabled[j] = true
					disabledCount++
					if disabledCount == len(pos) {
						fmt.Println("all disabled", disabled)
						goto part2
					}
				}
				hasVisited[j][p] = true
			}
		}

		// debug()
		// part1: stop condition
		allAtEnd := true
		for _, p := range pos {
			if !endIndexed[p] {
				allAtEnd = false
				break
			}
		}
		if allAtEnd {
			break
		}
		// move all
		instruction := instructions[i%len(instructions)]
		for j, p := range pos {
			if instruction == 'R' {
				pos[j] = graph[p][1]
			} else {
				pos[j] = graph[p][0]
			}
		}
		i++
	}
	fmt.Println("part1", i)

part2:
	for i, hist := range history {
		fmt.Printf("repeats i=%d %v\n", i, hist)
	}
	for t := history[0][0].T; true; t += history[0][1].T {
		ok := true
		for _, hist := range history {
			period := hist[1].T - hist[0].T
			offset := hist[0].T
			if (t-offset)%period != 0 {
				ok = false
				break
			}
		}
		if ok {
			fmt.Println("part2", t)
			break
		}
	}
}
