package main

import (
	"aoc/lib"
	"fmt"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()
	part1 := 0
	part2 := 0
	for _, line := range lib.Lines() {
		fmt.Println()
		fmt.Println("puzzle", line)
		p := problem{}
		p.parse(line, false)
		result := p.variations(0, 0)
		part1 += result
		fmt.Printf("%d\n", result)

		p.parse(line, true)
		result = p.variations(0, 0)
		part2 += result
		fmt.Printf("%d\n", result)
	}
	fmt.Printf("\npart1: %d\npart2: %d\n", part1, part2)
}

func noneWorking(str string) bool {
	for _, r := range str {
		if r == '.' {
			return false
		}
	}
	return true
}

type problem struct {
	template string
	chunks   []int
	cache    map[[2]int]int
}

func (p *problem) parse(line string, part2 bool) {
	parts := strings.Fields(line)
	template := parts[0]
	p.template = template
	p.chunks = lib.Map(strings.Split(parts[1], ","), lib.Int)
	p.cache = map[[2]int]int{}

	if part2 {
		templateRepeated := strings.Repeat(template+"?", 5)[0 : len(template)*5+4]
		chunksRepeated := lib.Map(strings.Split(strings.Repeat(parts[1]+",", 5)[0:len(parts[1])*5+4], ","), lib.Int)
		p.template = templateRepeated
		p.chunks = chunksRepeated
	}
}

func (p problem) variations(position int, chunkIdx int) (out int) {
	// memoization
	if v, hasCache := p.cache[[2]int{position, chunkIdx}]; hasCache {
		return v
	}
	defer func() {
		p.cache[[2]int{position, chunkIdx}] = out
	}()

	// main logic
	switch {
	case chunkIdx == len(p.chunks):
		if position > len(p.template) {
			position = len(p.template)
		}
		return lib.Ternary(strings.Count(p.template[position:], "#") == 0, 1, 0) // no more broken allowed
	case position >= len(p.template):
		return 0
	case p.template[position] == '#':
		if !p.canBeGroup(position, chunkIdx) {
			return 0
		}
		return p.variations(position+p.chunks[chunkIdx]+1, chunkIdx+1) // consume group
	case p.canBeGroup(position, chunkIdx):
		return 0 +
			p.variations(position+1, chunkIdx) + // skip one
			p.variations(position+p.chunks[chunkIdx]+1, chunkIdx+1) // consume group
	default: // .
		return p.variations(position+1, chunkIdx) // skip one
	}
}

func (p problem) canBeGroup(position int, chunkIdx int) bool {
	end := position + p.chunks[chunkIdx]
	if (position == 0 || p.template[position-1] != '#') &&
		end <= len(p.template) &&
		noneWorking(p.template[position:end]) &&
		(end == len(p.template) || p.template[end] != '#') {
		return true
	}
	return false
}
