package main

import (
	"aoc/lib"
	"fmt"
	"strings"
)

func main() {
	histories := lib.Map(lib.Lines(), func(line string) []int {
		return lib.Map(strings.Fields(line), lib.Int)
	})
	fmt.Println(lib.Sum(lib.Map(histories, setting{1}.Run)))
	fmt.Println(lib.Sum(lib.Map(histories, setting{2}.Run)))
}

type setting struct {
	part int
}

func (s setting) Run(history []int) (next int) {
	// fmt.Println()
	return s.run(history)
}

func (s setting) run(history []int) (next int) {
	// fmt.Println(history)
	// defer func() {
	// 	fmt.Println(history, next)
	// }()
	diff := make([]int, len(history)-1)
	var allZero bool = true
	for i := 0; i < len(history); i++ {
		if i < len(history)-1 {
			diff[i] = history[i+1] - history[i]
		}
		if history[i] != 0 {
			allZero = false
		}
	}
	if allZero {
		return 0
	}

	next = s.run(diff)
	if s.part == 1 {
		return history[len(history)-1] + next
	} else {
		return history[0] - next
	}
}
