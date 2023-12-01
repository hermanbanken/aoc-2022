package main

import (
	"aoc/lib"
	"bufio"
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	r := lib.Reader()
	defer r.Close()
	scanner := bufio.NewScanner(r)

	cycle := 1
	x := 1
	var values []int
	var crt []byte = bytes.Repeat([]byte{' '}, 40)
	check := func() {
		if cycle == 20 || (cycle-20)%40 == 0 {
			values = append(values, cycle*x)
		}
		if cycle > 0 && cycle%40 == 1 {
			fmt.Println(string(crt))
			crt = bytes.Repeat([]byte{' '}, 40)
		}
		crtPos := (cycle - 1) % 40
		if x-1 == crtPos || x == crtPos || x+1 == crtPos {
			crt[crtPos] = '#'
		}
	}

	for scanner.Scan() {
		t := scanner.Text()
		cmd := strings.Split(t, " ")
		switch cmd[0] {
		case "noop":
			cycle += 1
			check()
		case "addx":
			v, _ := strconv.Atoi(cmd[1])
			cycle += 1
			check()
			cycle += 1
			x += v
			check()
		}
	}

	fmt.Println(lib.Sum(values))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
