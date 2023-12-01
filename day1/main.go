package main

import (
	"aoc/lib"
	"log"
	"strings"
)

func main() {
	sum := 0
	lib.EachLine(func(line string) {
		if line == "" {
			return
		}
		log.Println(digitA(line), digitB(line))
		sum += digitA(line)*10 + digitB(line)
	})
	log.Println(sum)
}

var words = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func asDigit(line string, pos int) int {
	if line[pos] >= '0' && line[pos] <= '9' {
		return int(line[pos] - '0')
	}
	for i, w := range words {
		if strings.HasPrefix(line[pos:], w) {
			return i + 1
		}
	}
	return -1
}

func digitA(line string) int {
	for i := 0; i < len(line); i++ {
		d := asDigit(line, i)
		if d >= 0 {
			return d
		}
	}
	panic("no digit found in " + line)
}
func digitB(line string) int {
	for i := len(line) - 1; i >= 0; i-- {
		d := asDigit(line, i)
		if d >= 0 {
			return d
		}
	}
	panic("no digit found in " + line)
}
