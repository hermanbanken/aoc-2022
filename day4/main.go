package main

import (
	"aoc/lib"
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	r := lib.Reader()
	defer r.Close()

	var sum = 0
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		t := strings.TrimSpace(scanner.Text())
		if t == "" {
			break
		}

		sections := strings.Split(t, ",")
		a1, a2 := parseRange(sections[0])
		b1, b2 := parseRange(sections[1])
		if (a1 >= b1 && a2 <= b2) || (b1 >= a1 && b2 <= a2) {
			sum += 1
		}
	}
	fmt.Println("result:", sum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func parseRange(s string) (b int64, e int64) {
	parts := strings.Split(s, "-")
	b, _ = strconv.ParseInt(parts[0], 10, 64)
	e, _ = strconv.ParseInt(parts[1], 10, 64)
	return
}
