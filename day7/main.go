package main

import (
	"aoc/lib"
	"fmt"
	"slices"
	"sort"
	"strings"
)

var order = []string{
	"23456789TJQKA", // part 1
	"J23456789TQKA", // part 2
}
var question = 1

type Hand struct {
	hand  string
	cards [5]int
	bid   int
	cache [1]int
}

func (h Hand) String() string {
	return h.hand
}

func (h Hand) Options() []Hand {
	if !slices.Contains(h.cards[:], int('J')) {
		return []Hand{h}
	}
	return nil
}

func (h *Hand) Type() (out int) {
	if h.cache[0] != 0 {
		return h.cache[0]
	}
	defer func() {
		h.cache[0] = out
	}()

	sort.Ints(h.cards[:])
	var counts = map[int]int{}
	jokers := 0

	// part2:
	if question > 0 {
		for _, c := range h.cards {
			counts[c]++
			if c == int('J') {
				jokers++
			}
		}
	}

	countValues := [][2]int{}
	for key, count := range counts {
		countValues = append(countValues, [2]int{count, key})
	}
	sort.Slice(countValues, func(i, j int) bool {
		return countValues[i][0] > countValues[j][0]
	})

	// apply jokers
	if jokers > 0 {
		for idx := range countValues {
			if countValues[idx][1] != int('J') {
				// if h.hand == "JJJJ8" {
				// 	fmt.Println("bumping", string(rune(countValues[idx][1])), jokers)
				// }
				countValues[idx][0] += jokers
				jokers = 0
				for jidx := range countValues {
					if countValues[jidx][1] == int('J') {
						countValues[jidx][0] = 0
						break
					}
				}
				break
			}
		}
	}
	// resort after applying jokers
	sort.Slice(countValues, func(i, j int) bool {
		return countValues[i][0] > countValues[j][0]
	})

	if countValues[0][0] == 5 || countValues[0][0] == 4 {
		return countValues[0][0] + 1 // five/four of a kind
	}
	if len(countValues) >= 2 && countValues[0][0] == 3 && countValues[1][0] == 2 {
		return 4 // full house
	}
	if countValues[0][0] == 3 {
		return 3 // three of a kind
	}
	if len(countValues) >= 2 && countValues[0][0] == 2 && countValues[1][0] == 2 {
		return 2 // two pair
	}
	if countValues[0][0] == 2 {
		return 1 // one pair
	}
	return -1 // high card
}

func main() {
	in := lib.Lines()
	hands := lib.Map(in, func(s string) (out *Hand) {
		parts := strings.Fields(s)
		out = &Hand{hand: parts[0], bid: lib.Int(parts[1])}
		copy(out.cards[:], lib.Map([]byte(parts[0]), func(r byte) int { return int(r) }))
		return
	})
	sort.Sort(Hands(hands))
	winnings := 0
	for i, h := range hands {
		fmt.Println(h.hand, i+1, h.bid, h.Type())
		winnings += h.bid * (i + 1)
	}
	fmt.Println(winnings)
	// 248055134 high
	// 248860925 high
	// 247899149 ok
}

type Hands []*Hand

// Len implements sort.Interface.
func (h Hands) Len() int {
	return len(h)
}

// Less implements sort.Interface.
func (h Hands) Less(i int, j int) bool {
	a, b := h[i].Type(), h[j].Type()
	if a != b {
		return a < b
	}
	for idx := 0; idx < 5; idx++ {
		if h[i].hand[idx] != h[j].hand[idx] {
			return strings.IndexByte(order[question], h[i].hand[idx]) < strings.IndexByte(order[question], h[j].hand[idx])
		}
	}
	return false
}

// Swap implements sort.Interface.
func (h Hands) Swap(i int, j int) {
	tmp := h[i]
	h[i] = h[j]
	h[j] = tmp
}

var _ sort.Interface = Hands{}
