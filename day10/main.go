package main

import (
	"aoc/lib"
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type rope []lib.Coord

func (myrope rope) move(m lib.Coord) {
	myrope[0].Add(m)
	for i := 1; i < len(myrope); i++ {
		// follow
		d := myrope[i].Diff(myrope[i-1])
		if d < 2 || (d <= 2 && myrope[i].Diag(myrope[i-1])) {
			// nothing
		} else {
			// normal move & diagonal
			myrope[i].Add(myrope[i-1].Dir(myrope[i]))
		}
	}
}

func main() {
	r := lib.Reader()
	defer r.Close()
	scanner := bufio.NewScanner(r)

	cycle := 1
	x := 1
	var values []int
	check := func() {
		if cycle == 20 || (cycle-20)%40 == 0 {
			fmt.Println(cycle, x)
			values = append(values, cycle*x)
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
