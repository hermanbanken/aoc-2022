package lib

import (
	"log"
	"strconv"
)

type Coord struct {
	X, Y int
}

func (p *Coord) Add(o Coord) {
	p.X += o.X
	p.Y += o.Y
}
func (p Coord) Diff(o Coord) int {
	return AbsDiff(p.X, o.X) + AbsDiff(p.Y, o.Y)
}
func (p Coord) Diag(o Coord) bool {
	return AbsDiff(p.X, o.X) > 0 && AbsDiff(p.Y, o.Y) > 0
}
func (p Coord) Dir(o Coord) (r Coord) {
	if p.X > o.X {
		r.X = 1
	} else if o.X > p.X {
		r.X = -1
	}
	if p.Y > o.Y {
		r.Y = 1
	} else if o.Y > p.Y {
		r.Y = -1
	}
	return
}

func (p *Coord) Parse(dir string) {
	p.X = 0
	p.Y = 0
	switch dir {
	case "R":
		p.X = 1
		return
	case "U":
		p.Y = -1
		return
	case "L":
		p.X = -1
		return
	case "D":
		p.Y = 1
		return
	}
	log.Fatal("Unknown direction: ", dir)
}

func (p Coord) String() string {
	return strconv.Itoa(p.Y) + "," + strconv.Itoa(p.X)
}
