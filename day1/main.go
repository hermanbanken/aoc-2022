package main

import (
	"aoc/lib"
	"log"
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

func digitA(line string) int {
	for i := 0; i < len(line); i++ {
		if line[i] >= '0' && line[i] <= '9' {
			return int(line[i] - '0')
		}
	}
	panic("no digit found in " + line)
}
func digitB(line string) int {
	for i := len(line) - 1; i >= 0; i-- {
		if line[i] >= '0' && line[i] <= '9' {
			return int(line[i] - '0')
		}
	}
	panic("no digit found in " + line)
}
