package main

import (
	"aoc/lib"
	"log"
	"strings"
)

type Part struct {
	lib.Coord
	Length int
	Value  int
}

var nrPos = []Part{}

func main() {
	in := lib.Lines()
	for lineIdx, line := range in {
		for charIdx := 0; charIdx < len(line); charIdx++ {
			if isDigit(line[charIdx]) {
				end := strings.IndexFunc(line[charIdx:], notDigit)
				if end == -1 {
					end = len(line)
				}
				nrPos = append(nrPos, Part{Coord: lib.Coord{X: charIdx, Y: lineIdx}, Length: end - charIdx, Value: lib.Int(line[charIdx:end])})
				charIdx = end
				log.Println(nrPos[len(nrPos)-1])
			}
		}
	}
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func notDigit(char rune) bool {
	return char < '0' || char > '9'
}
