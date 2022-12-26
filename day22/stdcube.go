package main

import (
	"aoc/lib"
	"fmt"
	"math"
)

type action struct {
	side int
	rot  int
}

var (
	// standard cube moves & rotations
	right = map[int]action{1: {4, 1}, 2: {4, 0}, 3: {4, 3}, 4: {6, 2}, 5: {2, 0}, 6: {4, 2}}
	down  = map[int]action{1: {2, 0}, 2: {3, 0}, 3: {6, 0}, 4: {3, 1}, 5: {3, 3}, 6: {1, 0}}
	left  = map[int]action{1: {5, 3}, 2: {5, 0}, 3: {5, 1}, 4: {2, 0}, 5: {6, 2}, 6: {5, 2}}
	up    = map[int]action{1: {6, 0}, 2: {1, 0}, 3: {2, 0}, 4: {1, 3}, 5: {1, 1}, 6: {3, 0}}
)

// Side translates X:-1,Y:0 into 5
// Hardcoded all known input squares to their standard-cube face & rotation
func (c *Cube) Side(offset lib.Coord) (side int, facing int) {
	defer func() { side-- }()
	switch offset {
	case lib.Coord{X: -2, Y: 1}:
		return 6, -2
	case lib.Coord{X: -1, Y: 1}:
		return 5, 0
	case lib.Coord{X: -1, Y: 2}:
		return 5, -1
	case lib.Coord{X: -1, Y: 3}:
		return 6, 1
	case lib.Coord{X: 0, Y: 0}:
		return 1, 0
	case lib.Coord{X: 0, Y: 1}:
		return 2, 0
	case lib.Coord{X: 0, Y: 2}:
		return 3, 0
	case lib.Coord{X: 1, Y: 2}:
		return 4, 1
	case lib.Coord{X: 1, Y: 0}:
		return 4, -1
	}
	panic(fmt.Sprintf("Unknown offset %s", offset))
}

func (p Point) Next(facing int, dim int) (a Point, outFacing int) {
	var actions map[int]action
	switch facing {
	case 0:
		actions = right
	case 1:
		actions = down
	case 2:
		actions = left
	case 3:
		actions = up
	default:
		panic(fmt.Sprintf("No actions for facing %d", facing))
	}
	np := p.Coord.AddR(dir(facing)).ModR(dim)
	if p.AtEdge(facing, dim) {
		act := actions[p.Side+1]
		return Point{transp(np, act.rot, dim), act.side - 1}, (facing + act.rot) % 4
	}
	return Point{np, p.Side}, facing
}

func (p Point) AtEdge(facing int, dim int) bool {
	switch facing % 4 {
	case 0: // right
		return p.X == dim-1
	case 1: // down
		return p.Y == dim-1
	case 2: // left
		return p.X == 0
	case 3: // up
		return p.Y == 0
	default:
		panic(fmt.Sprintf("unknown facing %d", facing))
	}
}

// rotate N times 90° clockwise
func transpose(a lib.Coord, times int, dim int) lib.Coord {
	if times < 0 {
		times += 4
	}
	if times > 0 {
		a = rotate(a, lib.Coord{X: dim, Y: dim}, math.Pi/2)
		return transpose(a, times-1, dim)
		// a.X -= dim / 2
		// a.Y -= dim / 2
		// // 90° clockwise rotation: (x,y) becomes (y,-x)
		// tmp := a.X
		// a.X = a.Y
		// a.Y = -tmp
		// a.X += dim / 2
		// a.Y += dim / 2
		// return transpose(a, times-1, dim)
	}
	return a
}

// rotate N times 90° clockwise
func transp(a lib.Coord, times int, dim int) lib.Coord {
	if times < 0 {
		times += 4
	}
	if times > 0 {
		tmp := a.Y
		a.Y = a.X
		a.X = dim - tmp - 1

		// a.X -= dim / 2
		// a.Y -= dim / 2
		// // 90° clockwise rotation: (x,y) becomes (y,-x)
		// tmp := a.X
		// a.X = a.Y
		// a.Y = -tmp
		// a.X += dim / 2
		// a.Y += dim / 2
		return transp(a, times-1, dim)
	}
	return a
}
