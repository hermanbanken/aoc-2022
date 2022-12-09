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
	scanner := bufio.NewScanner(r)

	// Read stacks
	var hx,hy,tx,ty int
	var visited Grid = Grid{}

	var moves []string
	for scanner.Scan() {
		t := scanner.Text()
		moves = append(moves, t)
	}

	visited.Visit(tx, ty)
	for _, t := range moves {
		p := strings.Split(t, " ")
		direction := p[0]
		amount, _ := strconv.Atoi(p[1])
		fmt.Println(direction, amount)
		for i := 0; i < amount; i++ {
			fmt.Println(hx, hy, tx, ty, " diff ", diff(hx, tx)+ diff(hy, ty))
			mx,my := move(direction)
			hx += mx
			hy += my
	
			// follow
			d := diff(hx, tx) + diff(hy, ty)
			if d < 2 || (d <= 2 && diff(hx, tx) > 0 && diff(hy, ty) > 0) {
				// nothing
			} else {
				// normal move & diagonal
				tmx, tmy := dir(hx, hy, tx, ty)
				tx += tmx
				ty += tmy
			}
			visited.Visit(tx, ty)
		}
	}
	fmt.Println("last")
	fmt.Println(hx, hy, tx, ty, "diff", diff(hx, tx)+ diff(hy, ty))

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
		return a-b
	}
	return b-a
}

func (g *Grid) Visit(x,y int) {
	for g.data == nil {
		g.data = make(map[string]bool)
	}
	g.data[strconv.Itoa(y)+","+strconv.Itoa(x)] = true
	return
}
func (g Grid) Count() (count int) {
	for range g.data {
		count += 1
	}
	return
}

func move (dir string) (x, y int) {
	switch dir {
	case "R": return 1, 0
	case "U": return 0, -1
	case "L": return -1, 0
	case "D": return 0, 1
	}
	log.Fatal(dir)
	return 0, 0
}



func dir (ax, ay, bx, by int) (x, y int) {
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
