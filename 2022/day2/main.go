package main

import (
	"aoc/lib"
	"bufio"
	"fmt"
	"log"
	"strings"
)

func main() {
	r := lib.Reader()
	defer r.Close()

	var sum1 = 0
	var sum2 = 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		t := strings.TrimSpace(scanner.Text())
		sum1 += resultPart1(int(t[0]-'A'), int(t[2]-'X'))
		sum2 += resultPart2(int(t[0]-'A'), int(t[2]-'X'))
	}
	fmt.Println("result1:", sum1)
	fmt.Println("result2:", sum2)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func resultPart1(them int, us int) int {
	piece := us + 1
	if us == them {
		return 3 + piece
	}
	if us-1 == them || us+2 == them {
		return 6 + piece
	}
	return 0 + piece
}

func resultPart2(them int, result int) int {
	offset := result - 1
	pick := (them + offset + 3) % 3
	return result*3 + pick + 1
}
