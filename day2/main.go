package main

import (
	"bufio"
	"fmt"
	"lib"
	"log"
	"strings"
)

func main() {
	r := lib.Reader()
	defer r.Close()

	var sum = 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		t := strings.TrimSpace(scanner.Text())
		// sum += resultPart1(t[0], t[2]) + piecePart1(t[2])
		sum += resultPart2(t[0], t[2])
	}
	fmt.Println("result:", sum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func resultPart1(them byte, us byte) int {
	them -= 'A'
	us -= 'X'

	if us == them {
		return 3
	}
	if us-1 == them || us+2 == them {
		return 6
	}
	return 0
}

func piecePart1(us byte) int {
	return int(us - 'X' + 1)
}

func resultPart2(them byte, result byte) int {
	them -= 'A'
	result -= 'X'
	offset := int(result) - 1

	pick := (int(them) + offset) % 3
	if pick < 0 {
		pick += 3
	}
	return int(result)*3 + pick + 1
}
