package main

import (
	"aoc/lib"
	"fmt"
	"sort"
)

type Tile [2]int

func weight(c int) int {
	switch c {
	case int('#'):
		return 10
	case int('O'):
		return 20
	case int('.'):
		return 0
	}
	panic("invalid")
}

func (tile Tile) Less(other Tile) bool {
	if tile[0] == other[0] {
		return weight(tile[1]) > weight(other[1])
	}
	return tile[0] < other[0]
}

func main() {
	lines := lib.Lines()
	stacks := make([][]Tile, len(lines[0]))
	for _, line := range lib.Lines() {
		for x := range line {
			var last Tile
			if len(stacks[x]) > 0 {
				last = stacks[x][len(stacks[x])-1]
			}
			rank := last[0] + lib.Ternary(int(line[x]) == '#' || last[1] == '#', 1, 0)
			stacks[x] = append(stacks[x], Tile{rank, int(line[x])})
		}
	}
	// tilt north
	sum := 0
	for _, stack := range stacks {
		// fmt.Println(string(lib.Map(stack, func(t Tile) rune { return rune(t[1]) })))
		// fmt.Println(string(lib.Map(stack, func(t Tile) rune { return rune(t[0] + '0') })))
		sort.SliceStable(stack, func(i, j int) bool {
			return stack[i].Less(stack[j])
		})
		// fmt.Println(string(lib.Map(stack, func(t Tile) rune { return rune(t[1]) })))
		// fmt.Println(string(lib.Map(stack, func(t Tile) rune { return rune(t[0] + '0') })))
		// fmt.Println()
		stackSum := 0
		for y := range stack {
			if stack[y][1] == 'O' {
				stackSum += len(stack) - y
			}
		}
		fmt.Println(stackSum)
		sum += stackSum
	}
	fmt.Println("part1", sum)
}
