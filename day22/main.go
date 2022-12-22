package main

import (
	"aoc/lib"
	"fmt"
	"strings"
)

func main() {
	m := lib.InfinityMap[int]{}
	var instructions bool
	var steps []string
	var y = 0
	lib.EachLine(func(line string) {
		if line == "" {
			instructions = true
		}
		if instructions {
			for len(line) > 0 {
				idx := strings.IndexAny(line, "RL")
				if idx == -1 {
					steps = append(steps, line)
					line = ""
				} else {
					steps = append(steps, line[0:idx], line[idx:idx+1])
					line = line[idx+1:]
				}
			}
		} else {
			for x, c := range line {
				if c == ' ' {
					continue
				}
				m.Set(lib.Coord{X: x + 1, Y: y + 1}, lib.Ternary(c == '#', 2, 1))
			}
			y += 1
		}
	})

	fmt.Println(steps)
	fmt.Println(m.Draw(func(b int) byte { return byte(lib.Ternary(b > 0, lib.Ternary(b == 2, '#', '.'), ' ')) }))
}
