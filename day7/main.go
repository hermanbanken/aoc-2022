package main

import (
	"aoc/lib"
	"fmt"
	"sort"
	"strings"
)

type Hand struct {
	hand  string
	cards [5]int
	bid   int
}

func (h Hand) String() string {
	return h.hand
}

func (h Hand) Type() int {
	sort.Sort(sort.IntSlice(h.cards[:]))
	var counts = map[int]int{}
	for _, c := range h.cards {
		counts[c]++
	}
	countValues := []int{}
	for _, count := range counts {
		countValues = append(countValues, count)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(countValues)))
	if countValues[0] == 5 || countValues[0] == 4 {
		return countValues[0] + 1 // five/four of a kind
	}
	if len(countValues) >= 2 && countValues[0] == 3 && countValues[1] == 2 {
		return 4 // full house
	}
	if countValues[0] == 3 {
		return 3 // three of a kind
	}
	if len(countValues) >= 2 && countValues[0] == 2 && countValues[1] == 2 {
		return 2 // two pair
	}
	if countValues[0] == 2 {
		return 1 // one pair
	}
	return 0 // high card
}

func main() {
	in := lib.Lines()
	hands := lib.Map(in, func(s string) (out Hand) {
		parts := strings.Fields(s)
		out.hand = parts[0]
		out.cards[0] = int(byte(parts[0][0]))
		out.cards[1] = int(byte(parts[0][1]))
		out.cards[2] = int(byte(parts[0][2]))
		out.cards[3] = int(byte(parts[0][3]))
		out.cards[4] = int(byte(parts[0][4]))
		out.bid = lib.Int(parts[1])
		return
	})
	sort.Sort(Hands(hands))
	winnings := 0
	for i, h := range hands {
		fmt.Println(h.hand, i+1, h.bid)
		winnings += h.bid * (i + 1)
	}
	fmt.Println(winnings)
}

type Hands []Hand

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
			return strings.IndexByte(order, h[i].hand[idx]) < strings.IndexByte(order, h[j].hand[idx])
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

var order = "23456789TJQKA"
