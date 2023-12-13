package main

import (
	"aoc/lib"
	"fmt"
)

func main() {
	part1 := 0
	part2 := 0
	for count, puzzle := range puzzles(lib.Lines()) {
		sub1, sub2 := solve(count, puzzle)
		part1 += sub1
		part2 += sub2
		fmt.Println(sub1, sub2)
	}
	fmt.Println(part1, part2)
}

func puzzles(lines []string) (out [][]string) {
	start := 0
	end := 0
	for ; end < len(lines); end++ {
		if lines[end] == "" {
			out = append(out, lines[start:end])
			start = end + 1
		}
	}
	out = append(out, lines[start:])
	return
}

func transpose(lines []string) []string {
	var result []string
	for x := range lines[0] {
		var chars []byte
		for _, line := range lines {
			chars = append(chars, line[x])
		}
		result = append(result, string(chars))
	}
	return result
}

func solve(puzzle int, lines []string) (out int, part2 int) {
	lines_t := transpose(lines)
	hor := solveInner(puzzle, lines_t, false)
	ver := solveInner(puzzle, lines, false)
	defer func() {
		fmt.Printf("puzzle %d hor=%d ver=%d out=%d part2=%d\n", puzzle, hor, ver, out, part2)
		for _, l := range lines {
			// fmt.Println(l)
			_ = l
		}
		if hor == -1 && ver == -1 {
			panic("no mirror")
		}
		if out == part2 {
			panic("part1 & part2 must be different")
		}
	}()
	fmt.Println()
	out = lib.Max(hor, ver*100)

	// part2
	hor2 := solveInner(puzzle, lines_t, true)
	ver2 := solveInner(puzzle, lines, true)
	part2 = lib.Max(hor2, ver2*100)
	return
}

func solveInner(puzzle int, lines []string, slack bool) int {
outer:
	for i := range positions(len(lines)) {
		slackRemaining := lib.Ternary(slack, 1, 0)
		for p1, p2 := i, i+1; p1 >= 0 && p2 < len(lines); {
			if lines[p1] == lines[p2] {
				goto check
			} else if slackRemaining > 0 && dist(lines[p1], lines[p2]) == 1 {
				slackRemaining -= 1
				goto check
			} else {
				continue outer
			}
		check:
			if (p1 == 0 || p2 == len(lines)-1) && slackRemaining == 0 {
				return i + 1
			}
			p1 -= 1
			p2 += 1
		}
	}
	return -1
}

func positions(count int) (out []int) {
	var a = count/2 - 1 + count%2
	var b = count/2 + count%2
	for z := 0; z < count; z++ {
		if z%2 == 0 {
			out = append(out, a)
			a -= 1
		} else {
			out = append(out, b)
			b += 1
		}
	}
	return
}

func dist(a, b string) int {
	var result int
	for i := range a {
		if a[i] != b[i] {
			result += 1
		}
	}
	return result
}
