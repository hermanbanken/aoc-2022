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
	cardCounts := make([]int, len(cards))
	for i := range cards {
		cardCounts[i] = 1
	}
	matches := make([]int, len(cards))
	for i, card := range cards {
		card = strings.TrimPrefix(strings.TrimSpace(strings.TrimPrefix(card, "Card")), strconv.Itoa(i+1)+":")
		parts := strings.Split(card, "|")
		fmt.Println(parts)
		winners := strings.Fields(parts[0])
		have := strings.Fields(parts[1])
		overlap := len(lib.Intersect(lib.Map(winners, lib.Int), lib.Map(have, lib.Int)))
		fmt.Println("overlap", overlap)
		matches[i] = overlap
		if overlap > 0 {
			sum += 1 << (overlap - 1)
		}
	}
	fmt.Println("part1", sum)

	for i := range cards {
		for j := i + 1; j < i+1+matches[i]; j++ {
			fmt.Printf("adding %d cards of %d\n", cardCounts[i], i+1)
			cardCounts[j] += cardCounts[i]
		}
	}
	cardTotal := 0
	for _, count := range cardCounts {
		cardTotal += count
	}
	fmt.Println("part2", cardTotal)
}
