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
	// fmt.Println(lib.Sum(lib.Map(histories, part1head)))
	fmt.Println(lib.Sum(lib.Map(histories, part2head)))
}

func part1head(history []int) (next int) {
	fmt.Println()
	return part1(history)
}

func part1(history []int) (next int) {
	fmt.Println(history)
	defer func() {
		fmt.Println(history, next)
	}()
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

	next = part1(diff)
	return history[len(history)-1] + next
}

func part2head(history []int) (next int) {
	fmt.Println()
	return part2(history)
}
func part2(history []int) (next int) {
	fmt.Println(history)
	defer func() {
		fmt.Println(history, next)
	}()
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

	next = part2(diff)
	return history[0] - next
}
