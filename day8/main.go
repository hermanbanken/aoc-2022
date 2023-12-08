package main

import (
	"aoc/lib"
	"fmt"
	"strings"
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

	pos := nr("AAA")
	end := nr("ZZZ")
	i := 0
	for {
		if pos == end {
			break
		}
		if instructions[i%len(instructions)] == 'R' {
			pos = graph[pos][1]
		} else {
			pos = graph[pos][0]
		}
		i++
	}
	fmt.Println(nodesReverse[pos], i)
}
