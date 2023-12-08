package main

import (
	"aoc/lib"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	cards := lib.Lines()
	sum := 0
	for i, card := range cards {
		card = strings.TrimPrefix(strings.TrimSpace(strings.TrimPrefix(card, "Card")), strconv.Itoa(i+1)+":")
		parts := strings.Split(card, "|")
		fmt.Println(parts)
		winners := strings.Fields(parts[0])
		have := strings.Fields(parts[1])
		overlap := len(lib.Intersect(lib.Map(winners, lib.Int), lib.Map(have, lib.Int)))
		fmt.Println("overlap", overlap)
		if overlap > 0 {
			sum += 1 << (overlap - 1)
		}
	}
	fmt.Println("part1", sum)
}
