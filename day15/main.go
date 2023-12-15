package main

import (
	"aoc/lib"
	"fmt"
	"strings"
)

func main() {
	line := lib.Lines()[0]
	instructions := strings.Split(line, ",")
	var part1 = 0
	for _, inst := range instructions {
		r := hash(inst)
		fmt.Println(inst, r)
		part1 += r
	}
	fmt.Println("part1", part1)
}

func hash(str string) int {
	val := 0
	for i := range str {
		val += int(byte(str[i]))
		val *= 17
		val %= 256
	}
	return val
}
