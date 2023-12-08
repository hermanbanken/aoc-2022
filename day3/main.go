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

	fmt.Println("part1", lib.Sum(partsEngine))

	fmt.Println("part2")
	part2 := 0
	mp.EachCoord(func(c lib.Coord, r rune) bool {
		// ignore if not a gear
		if mp.GetOrDefault(c) != '*' {
			return true
		}

		digits := lib.Filter(c.Around(), func(c lib.Coord) bool {
			return !notDigit(mp.GetOrDefault(c))
		})
		fmt.Println("around", string(lib.Map(digits, mp.GetOrDefault)))

		gears := lib.UniqueUsingKey(lib.Map(lib.Map(digits, func(digit lib.Coord) string {
			width := 1
			// expand left
			for {
				if !notDigit(mp.GetOrDefault(digit.AddR(lib.Coord{X: -1}))) {
					digit = digit.AddR(lib.Coord{X: -1})
					width += 1
				} else {
					break
				}
			}
			// expand right
			for {
				if !notDigit(mp.GetOrDefault(digit.AddR(lib.Coord{X: width}))) {
					width += 1
				} else {
					break
				}
			}

			coords := []lib.Coord{digit}
			for i := 1; i < width; i++ {
				coords = append(coords, digit.AddR(lib.Coord{X: i}))
			}

			return string(lib.Map(coords, mp.GetOrDefault))
		}), lib.Int))

		if len(gears) == 2 {
			part2 += gears[0] * gears[1]
		}
		return true
	})
	fmt.Println("part2", part2)

}

func notDigit(char rune) bool {
	return char < '0' || char > '9'
}
