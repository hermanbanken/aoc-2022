package main

import (
	"aoc/lib"
	"fmt"
	"strconv"
)

type Part struct {
	lib.Coord
	Length int
	Value  int
}

var nrPos = []Part{}

func main() {
	mp := lib.InfinityMap[rune]{}
	mp.SetDefault('.')
	for y, line := range lib.Lines() {
		for x, char := range line {
			mp.Set(lib.Coord{Y: y, X: x}, char)
		}
	}

	fmt.Println(mp.Draw(func(r rune) byte {
		return byte(r)
	}))

	partsEngine := []int{}
	partsOther := []int{}
	mp.EachCoord(func(c lib.Coord, r rune) bool {
		// ignore if not a start of number
		left := mp.GetOrDefault(c.AddR(lib.Coord{X: -1}))
		if !notDigit(left) || notDigit(mp.GetOrDefault(c)) {
			return true
		}

		lookForSymbol := c.Around()
		word := []lib.Coord{c}
		width := 1
		for ; width < 100; width++ {
			rightCoord := c.AddR(lib.Coord{X: width})
			right := mp.GetOrDefault(rightCoord)
			if notDigit(right) {
				break
			}
			word = append(word, rightCoord)
			lookForSymbol = append(lookForSymbol, rightCoord.Around()...)
		}

		symbols := lib.Filter(lib.Map(lookForSymbol, mp.GetOrDefault), func(r rune) bool {
			return notDigit(r) && r != '.' && r != 0
		})
		hasSymbol := len(symbols) > 0

		data := string(lib.Map(word, mp.GetOrDefault))
		fmt.Println(string(data), string(symbols), len(symbols), symbols, hasSymbol)
		part, err := strconv.Atoi(data)
		if err != nil {
			fmt.Println("error", err, word, data)
		}
		if hasSymbol {
			partsEngine = append(partsEngine, part)
		} else {
			partsOther = append(partsOther, part)
		}
		return true
	})

	fmt.Println("part1", lib.Sum(partsEngine), partsEngine, lib.Sum(partsOther), partsOther)
}

func notDigit(char rune) bool {
	return char < '0' || char > '9'
}
