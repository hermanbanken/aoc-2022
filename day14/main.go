package main

import (
	"aoc/lib"
	"fmt"
	"log"
	"strings"
)

func main() {
	var maxx, maxy, minx, miny = 0, 0, 10000000, 10000000 //  573 158 489 13
	var paths [][]lib.Coord = nil
	var sand = lib.Coord{X: 500, Y: 0}
	lib.EachLine(func(line string) {
		path := []lib.Coord{}
		for _, coord := range strings.Split(line, " -> ") {
			xy := strings.Split(coord, ",")
			c := lib.Coord{X: lib.Int(xy[0]), Y: lib.Int(xy[1])}
			if maxx < c.X {
				maxx = c.X
			}
			if maxy < c.Y {
				maxy = c.Y
			}
			if minx > c.X {
				minx = c.X
			}
			if miny > c.Y {
				miny = c.Y
			}
			path = append(path, c)
		}
		paths = append(paths, path)
	})
	fmt.Println(paths, maxx, maxy, minx, miny)

	var filled []lib.Coord = nil

	for {
		particle := Particle{sand}
		for particle.Y <= maxy && particle.Down(paths, filled) {
			// log.Println("at", particle)
			// time.Sleep(10 * time.Millisecond)
		}
		if particle.Y > maxy {
			log.Println("done", len(filled))
			break
		}
		// log.Println("deposited at", particle)
		filled = append(filled, particle.Coord)
		// time.Sleep(1 * time.Second)
	}

}

type Particle struct {
	lib.Coord
}

func (p *Particle) Down(paths [][]lib.Coord, filled []lib.Coord) bool {
	var options = []lib.Coord{
		{X: p.X, Y: p.Y + 1},
		{X: p.X - 1, Y: p.Y + 1},
		{X: p.X + 1, Y: p.Y + 1},
	}
options:
	for len(options) > 0 {
		for _, path := range paths {
			if intersect(path, options[0]) {
				options = options[1:]
				continue options
			}
		}
		for _, f := range filled {
			if f == options[0] {
				options = options[1:]
				continue options
			}
		}
		p.X = options[0].X
		p.Y = options[0].Y
		return true
	}
	return false
}

func intersect(path []lib.Coord, c lib.Coord) bool {
	for i := 0; i < len(path)-1; i++ {
		a := path[i]
		b := path[i+1]
		dx := b.X - a.X
		dy := b.Y - a.Y
		if dx == 0 {
			if c.X == b.X && within(a.Y, b.Y, c.Y) {
				return true
			}
			continue
		} else if dy == 0 {
			if c.Y == b.Y && within(a.X, b.X, c.X) {
				return true
			}
			continue
		}
		panic(fmt.Sprintf("non vertical/horizontal line: %s", path))
	}
	return false
}

func within(a, b int, p int) bool {
	if b > a {
		return p >= a && p <= b
	}
	return p >= b && p <= a
}
