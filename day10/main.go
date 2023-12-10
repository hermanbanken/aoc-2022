/*
| is a vertical pipe connecting north and south.
- is a horizontal pipe connecting east and west.
L is a 90-degree bend connecting north and east.
J is a 90-degree bend connecting north and west.
7 is a 90-degree bend connecting south and west.
F is a 90-degree bend connecting south and east.
. is ground; there is no pipe in this tile.
S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
*/
package main

import (
	"aoc/lib"
	"fmt"
)

func main() {
	lines := lib.Lines()
	_ = lines
	start := lib.Coord{X: 0, Y: 0}
	mp := lib.InfinityMap[rune]{}
	mp.SetDefault('.')
	for y, line := range lines {
		for x, c := range line {
			if c == 'S' {
				start = lib.Coord{X: x, Y: y}
			}
			mp.Set(lib.Coord{X: x, Y: y}, c)
		}
	}

	isLoop := lib.InfinityMap[bool]{}
	isLoop.Set(start, true)

	// part 1
	dir := lib.Coord{X: 0, Y: 0}.CrossAround()
	pos := []lib.Coord{start, start, start, start}
	steps := 1
	for {
		var prune []int
		for i := 0; i < len(dir); i++ {
			next := pos[i].AddR(dir[i])
			shape, _ := mp.Get(next)
			if isInverted(dir[i], shapes[shape][0]) {
				dir[i] = shapes[shape][1]
				pos[i] = next
			} else if isInverted(dir[i], shapes[shape][1]) {
				dir[i] = shapes[shape][0]
				pos[i] = next
			} else {
				prune = append(prune, i) // remove not connected from start position
				fmt.Println("prune", pos[i], dir[i], string(shape))
			}
		}

		for i := len(prune) - 1; i >= 0; i-- {
			dir = append(dir[:prune[i]], dir[prune[i]+1:]...)
			pos = append(pos[:prune[i]], pos[prune[i]+1:]...)
		}

		for i := 0; i < len(pos); i++ {
			isLoop.Set(pos[i], true)
		}
		if pos[0] == pos[1] {
			fmt.Println("part1, found half-way at", pos[0], "steps", steps)
			break
		}
		steps += 1
	}

	// part 2
	scanLine := lib.Coord{X: -1, Y: 0}.Path([]struct {
		Dir   lib.Coord
		Count int
	}{
		{Dir: lib.Coord{Y: 1}, Count: len(lines)},
	}...)
	crossings := make([]int, len(scanLine))
	inside := lib.InfinityMap[bool]{}

	for x := 0; x < len(lines[0]); x++ {
		for i, c := range scanLine {

			shape, _ := mp.Get(c)
			if isLoop.GetOrDefault(c) {
				switch shape {
				case '|', '7', 'F':
					crossings[i] += 1
				}
			}
			if !isLoop.GetOrDefault(c) && crossings[i]%2 == 1 {
				inside.Set(c, true)
			}
			scanLine[i] = c.AddR(lib.Coord{X: 1})
		}
	}

	fmt.Println("part2", inside.Len())
	fmt.Println(inside.Draw(func(b bool) byte {
		if b {
			return '#'
		}
		return '.'
	}))
}

func isInverted(a, b lib.Coord) bool {
	return a.X == -b.X && a.Y == -b.Y
}

var shapes = map[rune][2]lib.Coord{
	'|': {lib.Coord{Y: 1}, lib.Coord{Y: -1}},
	'-': {lib.Coord{X: -1}, lib.Coord{X: 1}},
	'L': {lib.Coord{Y: -1}, lib.Coord{X: 1}},
	'J': {lib.Coord{Y: -1}, lib.Coord{X: -1}},
	'7': {lib.Coord{Y: 1}, lib.Coord{X: -1}},
	'F': {lib.Coord{Y: 1}, lib.Coord{X: 1}},
}
