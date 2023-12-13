package main

import (
	"aoc/lib"
	"fmt"
	"log"
)

func main() {
	log.Println(positions(10), positions(9), positions(2))

	//405
	lines := lib.Lines()
	start := 0
	end := 0
	sum := 0
	count := 0
	for ; end < len(lines); end++ {
		if lines[end] == "" {
			result := solve(count, lines[start:end])
			sum += result
			fmt.Println(result)
			start = end
			count += 1
		}
	}
	result := solve(count, lines[start+1:])
	sum += result
	fmt.Println(result)
	fmt.Println(sum)
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

func solve(puzzle int, lines []string) (out int) {
	hor := solveInner(transpose(lines))
	ver := solveInner(lines)
	defer func() {
		fmt.Printf("puzzle %d hor=%d ver=%d out=%d\n", puzzle, hor, ver, out)
		for _, l := range lines {
			fmt.Println(l)
			_ = l
		}
	}()
	fmt.Println()
	if hor == -1 && ver == -1 {
		panic("no mirror")
	}
	if hor > ver {
		return hor
	} else {
		return ver * 100
	}
}

func solveInner(lines []string) int {
outer:
	for i := range positions(len(lines)) {
		for p1, p2 := i, i+1; p1 >= 0 && p2 < len(lines); {
			if lines[p1] == lines[p2] {
				if p1 == 0 || p2 == len(lines)-1 {
					return i + 1
				}
				p1 -= 1
				p2 += 1
			} else {
				continue outer
			}
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
