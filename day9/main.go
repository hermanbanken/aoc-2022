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

	// Read stacks
	var myrope rope = make([]lib.Coord, 10)
	var visited = lib.Set{}

	visited.Add(lib.Last(myrope))
	for scanner.Scan() {
		t := scanner.Text()
		p := strings.Split(t, " ")
		amount, _ := strconv.Atoi(p[1])
		var direction lib.Coord
		direction.Parse(p[0])
		fmt.Println(p[0], amount)
		for amount > 0 {
			amount--
			fmt.Println(myrope)
			myrope.move(direction)
			visited.Add(lib.Last(myrope))
		}
	}
	fmt.Println(visited.Size())

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
