package main

import (
	"aoc/lib"
	"fmt"
	"log"
	"strings"
)

type M struct {
	lib.InfinityMap[int8]
	side   byte
	facing int
	pos    lib.Coord
	Dim    int
	Fold   string
}

func (m M) nextPos() lib.Coord {
	// straight
	v, has := m.Get(m.pos.AddR(dir(m.facing)))
	if has {
		return m.pos.AddR(dir(m.facing))
	}

	// TODO diagonal => rotate 1
	// TODO horse-jump => rotate 2
	// TODO horse-jump + diagonal

	round := lib.Coord{X: m.pos.X / m.Dim, Y: m.pos.Y / m.Dim}
	modv := lib.Coord{X: m.pos.X % m.Dim, Y: m.pos.Y % m.Dim}
	flip := lib.Coord{X: m.Dim - m.pos.X%m.Dim, Y: m.Dim - m.pos.Y%m.Dim}

	// diagonal walk
	mod(m.pos.AddR(dir(m.facing-1).MultR(m.Dim)), m.Dim) // left point
	mod(m.pos.AddR(dir(m.facing+1).MultR(m.Dim)), m.Dim) // right point
	dir1 := dir(m.facing).AddR(dir(m.facing - 1))
	dir2 := dir(m.facing).AddR(dir(m.facing + 1))

	// the base point of blocks diagonal
	d1 := round.AddR(dir(m.facing).MultR(m.Dim)).AddR(dir(m.facing - 1).MultR(m.Dim)).AddR(flip)
	d2 := round.AddR(dir(m.facing).MultR(m.Dim)).AddR(dir(m.facing + 1).MultR(m.Dim)).AddR(flip)
	if _, hasD1 := m.Get(d1); hasD1 {
		m.pos.AddR(dir(m.facing))
	}

	if _, has = m.Get(p); has { /**/
	}
}

func mod(c lib.Coord, dim int) lib.Coord {
	return lib.Coord{X: c.X % dim, Y: c.Y % dim}
}

var fold0 = strings.Trim(`
  A 
DCE 
  FB
`, "\n\r")

var fold1 = strings.Trim(`
 AB
 E 
CF 
D  
`, "\n\r")

// _ = fold0
// _ = fold1
/*
..D..
.CAB.
..E..
..F..
*/
func (m M) moveRotate() (facing int) {
	rows := strings.Split(m.Fold, "\n")
	_ = rows
	strings.IndexByte(m.Fold, m.side)
	return 0
}
func (m M) moveSide() (dstSide byte) {
	switch m.side {
	case 'A':
		return "BECD"[m.facing]
	case 'B':
		return "CEAD"[m.facing]
	case 'C':
		return "AEBD"[m.facing]
	case 'D':
		return "BACF"[m.facing]
	case 'E':
		return "BFCA"[m.facing]
	case 'F':
		return "BDCE"[m.facing]
	default:
		panic("invalid side")
	}
}

func main() {
	m := &M{}
	m.SetDefault(' ')
	var instructions bool
	var steps []string
	var y = 0
	var first = true
	lib.EachLine(func(line string) {
		if line == "" {
			instructions = true
		}
		if instructions {
			for len(line) > 0 {
				idx := strings.IndexAny(line, "RL")
				if idx == -1 {
					steps = append(steps, line)
					line = ""
				} else {
					steps = append(steps, line[0:idx], line[idx:idx+1])
					line = line[idx+1:]
				}
			}
		} else {
			if m.Dim == 0 {
				m.Dim = (lib.Ternary(len(line) == 150, 50, 4))
				m.Fold = (lib.Ternary(len(line) == 150, fold1, fold0))
			}
			for x, c := range line {
				if c == ' ' {
					continue
				}
				if first {
					first = false
					m.pos.X = x
					m.pos.Y = y
					m.facing = 0
				}
				m.Set(lib.Coord{X: x, Y: y}, int8(c))
			}
			y += 1
		}
	})

	fmt.Println(m.Dim, steps)
	// fmt.Println(m.Draw(func(b int8) byte { return byte(b) }))

	fmt.Println(m.pos)
	m.simulate(steps)
	fmt.Println("part1", m.pos, 1000*m.pos.Y+4*m.pos.X+m.facing)

}

func (m *M) simulate(path []string) {
	if len(path) == 0 {
		return
	}
	fmt.Println(path[0])
	switch path[0] {
	case "R":
		m.facing = (m.facing + 1) % 4
	case "L":
		m.facing = (m.facing + 3) % 4
	default:
		for i := lib.Int(path[0]); i > 0; i-- {
			n := m.next()
			if n == m.pos {
				break
			}
			m.pos = n
			fmt.Println(m.pos)
		}
	}
	m.simulate(path[1:])
}

func (m *M) next() lib.Coord {
	p := m.pos
	p.Add(dir(m.facing))
	v, has := m.Get(p)
	if has && v == '#' {
		return m.pos
	}
	if !has {
		switch m.facing {
		case 0:
			for x := m.Bounds()[0].X; x < m.Bounds()[1].X; x++ {
				if v, has := m.Get(lib.Coord{Y: m.pos.Y, X: x}); has {
					return lib.Ternary(v == '.', lib.Coord{Y: m.pos.Y, X: x}, m.pos)
				}
			}
		case 1:
			for y := m.Bounds()[0].Y; y < m.Bounds()[1].Y; y++ {
				if v, has := m.Get(lib.Coord{X: m.pos.X, Y: y}); has {
					return lib.Ternary(v == '.', lib.Coord{X: m.pos.X, Y: y}, m.pos)
				}
			}
		case 2:
			for x := m.Bounds()[1].X; x > m.Bounds()[0].X; x-- {
				if v, has := m.Get(lib.Coord{Y: m.pos.Y, X: x}); has {
					return lib.Ternary(v == '.', lib.Coord{Y: m.pos.Y, X: x}, m.pos)
				}
			}
		case 3:
			for y := m.Bounds()[1].Y; y > m.Bounds()[0].Y; y-- {
				if v, has := m.Get(lib.Coord{X: m.pos.X, Y: y}); has {
					return lib.Ternary(v == '.', lib.Coord{X: m.pos.X, Y: y}, m.pos)
				}
			}
		default:
			panic("wrong facing")
		}
		return m.pos
	}
	return p
}
func dir(facing int) (dir lib.Coord) {
	switch facing {
	case 0:
		dir = lib.Coord{X: 1, Y: 0}
	case 1:
		dir = lib.Coord{X: 0, Y: 1}
	case 2:
		dir = lib.Coord{X: -1, Y: 0}
	case 3:
		dir = lib.Coord{X: 0, Y: -1}
	default:
		log.Fatalf("invalid facing %d", facing)
	}
	return
}
