package main

import (
	"aoc/lib"
	"fmt"
)

func main() {
	stars := make([]lib.Coord, 0)

	lines := lib.Lines()
	_ = lines
	mp := lib.InfinityMap[rune]{}
	mp.SetDefault('.')
	for y, line := range lines {
		for x, c := range line {
			mp.Set(lib.Coord{X: x, Y: y}, c)
			if c == '#' {
				stars = append(stars, lib.Coord{X: x, Y: y})
			}
		}
	}

	// find empty rows and columns
	var emptyRows = make([]bool, len(lines))
	var emptyCols = make([]bool, len(lines[0]))
	for x := 0; x < len(lines[0]); x++ {
		something := false
		for y := 0; y < len(lines); y++ {
			if d, _ := mp.Get(lib.Coord{X: x, Y: y}); d == '#' {
				something = true
				break
			}
		}
		emptyCols[x] = !something
	}
	for y := 0; y < len(lines); y++ {
		something := false
		for x := 0; x < len(lines[0]); x++ {
			if d, _ := mp.Get(lib.Coord{X: x, Y: y}); d == '#' {
				something = true
				break
			}
		}
		emptyRows[y] = !something
	}
	fmt.Println("empty rows", emptyRows)
	fmt.Println("empty cols", emptyCols)

	bend := 1_000_000 - 1 // 0_000 //for part2

	dist := func(a, b lib.Coord) (weight int) {
		dir := lib.Coord{X: b.X - a.X, Y: b.Y - a.Y}
		// fmt.Println(a, b, dir)
		for i := 0; a.X+i != b.X; i += lib.Unit(dir.X) {
			if emptyCols[a.X+i] {
				weight += bend
			}
		}
		for i := 0; a.Y+i != b.Y; i += lib.Unit(dir.Y) {
			if emptyRows[a.Y+i] {
				weight += bend
			}
		}
		return weight + lib.AbsDiff(a.X, b.X) + lib.AbsDiff(a.Y, b.Y)
	}

	g := lib.NewGraph()
	for i := range stars {
		for j := range stars {
			if i == j {
				continue
			}
			weight := dist(stars[i], stars[j])
			g.Add(i, j, weight)
		}
	}
	out, _ := lib.FloydWarshall(g)

	var sum = 0
	var seenPair = make(map[lib.Coord]bool)
	for i := range stars {
		for j := range stars {
			pair := lib.Coord{X: i, Y: j}
			if j < i {
				pair = lib.Coord{X: j, Y: i}
			}
			if i == j || seenPair[pair] {
				continue
			}
			seenPair[pair] = true
			sum += out[lib.Vertex(i)][lib.Vertex(j)]
		}
	}

	fmt.Println("part2", sum)
}
