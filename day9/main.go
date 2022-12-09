package main

import (
	"aoc/lib"
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type pos struct {
	x, y int
}

func (p *pos) add(o pos) {
	p.x += o.x
	p.y += o.y
}
func (p pos) diff(o pos) int {
	return diff(p.x, o.x) + diff(p.y, o.y)
}
func (p pos) diag(o pos) bool {
	return diff(p.x, o.x) > 0 && diff(p.y, o.y) > 0
}
func (p pos) dir(o pos) (r pos) {
	if p.x > o.x {
		r.x = 1
	} else if o.x > p.x {
		r.x = -1
	}
	if p.y > o.y {
		r.y = 1
	} else if o.y > p.y {
		r.y = -1
	}
	return
}

type rope []pos

func (myrope rope) move(direction string) {
	m := move(direction)
	myrope[0].add(m)
	for i := 1; i < len(myrope); i++ {
		// follow
		d := myrope[i].diff(myrope[i-1])
		if d < 2 || (d <= 2 && myrope[i].diag(myrope[i-1])) {
			// nothing
		} else {
			// normal move & diagonal
			myrope[i].add(myrope[i-1].dir(myrope[i]))
		}
	}
}

func main() {
	r := lib.Reader()
	defer r.Close()
	scanner := bufio.NewScanner(r)

	// Read stacks
	var myrope rope = make([]pos, 10)
	var visited Grid = Grid{}

	var moves []string
	for scanner.Scan() {
		t := scanner.Text()
		moves = append(moves, t)
	}

	visited.Visit(myrope[9])
	for _, t := range moves {
		p := strings.Split(t, " ")
		direction := p[0]
		amount, _ := strconv.Atoi(p[1])
		fmt.Println(direction, amount)
		for i := 0; i < amount; i++ {
			fmt.Println(myrope)
			myrope.move(direction)
			visited.Visit(myrope[9])
		}
	}
	fmt.Println(visited.Count())

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

type Grid struct {
	data map[string]bool
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func (g *Grid) Visit(p pos) {
	for g.data == nil {
		g.data = make(map[string]bool)
	}
	g.data[strconv.Itoa(p.y)+","+strconv.Itoa(p.x)] = true
	return
}
func (g Grid) Count() (count int) {
	for range g.data {
		count += 1
	}
	return
}

func move(dir string) pos {
	switch dir {
	case "R":
		return pos{1, 0}
	case "U":
		return pos{0, -1}
	case "L":
		return pos{-1, 0}
	case "D":
		return pos{0, 1}
	}
	log.Fatal(dir)
	return pos{0, 0}
}

func dir(ax, ay, bx, by int) (x, y int) {
	if ax > bx {
		x = 1
	} else if bx > ax {
		x = -1
	}
	if ay > by {
		y = 1
	} else if by > ay {
		y = -1
	}
	return
}
