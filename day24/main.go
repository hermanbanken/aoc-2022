package main

import (
	"aoc/lib"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Blizzard struct {
	pos int
	dir int
}

func main() {
	lines := lib.Lines()
	g := lib.Grid{H: len(lines), W: len(lines[0])}
	cells := strings.Join(lines, "")
	blizz := parseBlizzards(g, cells)
	bm := blizzMap(blizz)

	start := 1
	end := len(cells) - 2
	t := 0

	empty := func(p int) bool {
		if p == start || p == end {
			return true
		}
		if p%g.W == 0 || p%g.W == g.W-1 {
			return false
		}
		if p/g.W == 0 || p/g.W == g.H-1 {
			return false
		}
		_, hasBlizzard := bm[p]
		return !hasBlizzard
	}

	moveBlizzard := func(b Blizzard) Blizzard {
		np := b.pos + b.dir
		if np%g.W == 0 {
			np += g.W - 2
		}
		if np%g.W == g.W-1 {
			np -= g.W - 2
		}
		if np/g.W == 0 {
			np += g.W * (g.H - 2)
		}
		if np/g.W == g.H-1 {
			np -= g.W * (g.H - 2)
		}
		return Blizzard{np, b.dir}
	}

	g.CanMove = func(posA, posB int) bool { return empty(posB) }

	cycle := (g.W - 2) * (g.H - 2)
	fmt.Println("cycle", cycle)

	pos := []int{start}
	minTime := map[int]int{}
	trip := 0
	for ; t < 1000; t++ {
		fmt.Println("Minute", t+1, len(pos))
		blizz = lib.Map(blizz, moveBlizzard)
		bm = blizzMap(blizz)

		newPos := []int{}
		for _, p := range pos {
			if empty(p) {
				newPos = append(newPos, p)
			}
			newPos = append(newPos, g.Moves(p)...)
		}
		remainingPos := []int{}
		for _, np := range newPos {
			if t > 0 && (minTime[np] < t-cycle || minTime[np] == t) {
				continue
			}
			remainingPos = append(remainingPos, np)
			minTime[np] = t
		}
		sort.Ints(remainingPos)
		pos = remainingPos

		// Detection
		if trip == 2 && pos[len(pos)-1] == end {
			fmt.Printf("part 2 / at end in %d minutes\n", t+1)
			visualize(g, bm, []int{end})
			os.Exit(0)
		}
		if trip == 1 && pos[0] == start {
			trip += 1
			fmt.Printf("at start in %d minutes\n", t+1)
			pos = []int{start}
			visualize(g, bm, pos)
		}
		if trip == 0 && pos[len(pos)-1] == end {
			trip += 1
			fmt.Printf("part 1 / at end in %d minutes\n", t+1)
			pos = []int{end}
			visualize(g, bm, pos)
		}

		if cycle == 25 {
			visualize(g, bm, remainingPos)
		}
	}
}

func parseBlizzards(g lib.Grid, cells string) (out []Blizzard) {
	for i, char := range cells {
		switch char {
		case '>':
			out = append(out, Blizzard{i, 1})
		case '<':
			out = append(out, Blizzard{i, -1})
		case '^':
			out = append(out, Blizzard{i, -g.W})
		case 'v':
			out = append(out, Blizzard{i, g.W})
		}
	}
	return
}

func blizzMap(bs []Blizzard) (out map[int][]Blizzard) {
	out = map[int][]Blizzard{}
	for _, b := range bs {
		out[b.pos] = append(out[b.pos], b)
	}
	return out
}

func visualize(g lib.Grid, bm map[int][]Blizzard, pos []int) {
	exp := map[int]bool{}
	for _, p := range pos {
		exp[p] = true
	}

	sb := strings.Builder{}
	for y := 0; y < g.H; y++ {
		for x := 0; x < g.W; x++ {
			p := y*g.W + x
			c := len(bm[p])
			if (x == 0 || x == g.W-1 || y == 0 || y == g.H-1) && p != 1 && p != g.W*g.H-2 {
				sb.WriteRune('#')
			} else if c > 1 {
				sb.WriteString(strconv.Itoa(c))
			} else if c == 1 {
				sb.WriteRune(char(bm[p][0].dir))
			} else if exp[p] {
				sb.WriteRune('E')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}
	fmt.Println(sb.String())
}

func char(dir int) rune {
	if dir == -1 {
		return '<'
	}
	if dir == 1 {
		return '>'
	}
	if dir < 0 {
		return '^'
	}
	return 'v'
}
