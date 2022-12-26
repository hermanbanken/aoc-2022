package main

import (
	"aoc/lib"
	"fmt"
	"log"
	"math"
	"strings"
)

type Point struct {
	lib.Coord
	Side int
}

func (p Point) String() string {
	return fmt.Sprintf("%d,%d:%d", p.X, p.Y, p.Side+1)
}

type Cube struct {
	sides       [6]map[Point]lib.Coord
	sideFacings [6]int
	data        lib.InfinityMap[int8]
	pos         Point
	facing      int
	Dim         int
}

func (c *Cube) PositionFacing() (p lib.Coord, facing int) {
	p = c.sides[c.pos.Side][c.pos]
	facing = (c.sideFacings[c.pos.Side] + c.facing + 4) % 4
	return
}

func (c *Cube) Set(offset, p lib.Coord, char int8) {
	c.data.Set(p, char)
	sidePos := p.AddR(offset).ModR(c.Dim)
	sideCoord := p.AddR(offset).AddR(sidePos.MultR(-1)).DivR(c.Dim)

	side, facing := c.Side(sideCoord)
	c.sideFacings[side] = facing
	if c.sides[side] == nil {
		c.sides[side] = make(map[Point]lib.Coord)
	}
	cp := transp(sidePos, (4-facing)%4, c.Dim)
	if cp.X >= c.Dim || cp.X < 0 || cp.Y >= c.Dim || cp.Y < 0 {
		panic(fmt.Sprintf("invariant %s from %s (rot: %d)", cp, sidePos, (4-facing)%4))
	}
	c.sides[side][Point{cp, side}] = p
}

func (c *Cube) Get(a Point) (int8, bool) {
	p, hasP := c.sides[a.Side][a]
	if !hasP {
		panic(fmt.Sprintf("No position %s on side %d: \n%v", a.Coord, a.Side+1, c.sides[a.Side]))
	}
	return c.data.Get(p)
}

func (c *Cube) Next() (a Point, facing int, v int8) {
	np, nf := c.pos.Next(c.facing, c.Dim)
	v, _ = c.Get(np)
	return np, nf, v
}

type M struct {
	lib.InfinityMap[int8]
	facing int
	pos    lib.Coord
	Dim    int
}

type Stance struct {
	lib.Coord
	facing int
}

func read() (m *M, cube *Cube, steps []string) {
	m = &M{}
	m.SetDefault(' ')
	cube = &Cube{sides: [6]map[Point]lib.Coord{}, data: lib.InfinityMap[int8]{}}
	var instructions bool
	var y = 0
	var first = true
	var offset lib.Coord
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
				cube.Dim = (lib.Ternary(len(line) == 150, 50, 4))
			}
			for x, c := range line {
				if c == ' ' {
					continue
				}
				p := lib.Coord{X: x, Y: y}
				if first {
					first = false
					m.pos = p
					m.facing = 0
					offset = p.MultR(-1)
				}
				cube.Set(offset, p, int8(c))
				m.Set(p, int8(c))
			}
			y += 1
		}
	})
	return m, cube, steps
}

func rotate(p lib.Coord, origin lib.Coord, angle float64) lib.Coord {
	var s = int(math.Sin(angle))
	var c = int(math.Cos(angle))

	// translate point back to origin:
	p.X -= origin.X
	p.Y -= origin.Y

	// rotate point
	var xnew = p.X*c - p.Y*s
	var ynew = p.X*s + p.Y*c

	// translate point back:
	p.X = xnew + origin.X
	p.Y = ynew + origin.Y
	return p
}

func main() {
	m, _, steps := read()
	fmt.Println("drawing:", m.Bounds())
	fmt.Println(m.Draw(func(i int8) byte { return byte(i) }))
	fmt.Println(m.pos)
	m.simulate(steps)
	part1 := m.pos.AddR(lib.Coord{X: 1, Y: 1})
	fmt.Println("part1", part1, 1000*(part1.Y)+4*(part1.X)+m.facing)

	fmt.Println("\npart2 starting")
	m, c, steps := read()
	c.simulate(steps, m)
	part2, facing := c.PositionFacing()
	part2 = part2.AddR(lib.Coord{X: 1, Y: 1})
	fmt.Println(m.Draw(func(i int8) byte { return byte(i) }))
	fmt.Println("part2", part2, facing, 1000*(part2.Y)+4*(part2.X)+facing)
}

func (m *M) simulate(path []string) {
	if len(path) == 0 {
		return
	}
	// fmt.Println(path[0])
	switch path[0] {
	case "R":
		m.facing = (m.facing + 1) % 4
	case "L":
		m.facing = (m.facing + 3) % 4
	default:
		for i := lib.Int(path[0]); i > 0; i-- {
			var n Stance = Stance{m.pos, m.facing}
			n = Stance{m.nextPart1(), m.facing}
			if n.Coord == m.pos {
				// fmt.Println(m.pos, m.facing, "not moved")
				break
			}
			m.pos = n.Coord
			m.facing = n.facing
			// fmt.Println(m.pos, m.facing)
		}
	}
	m.simulate(path[1:])
}

func (c *Cube) simulate(path []string, m *M) {
	if len(path) == 0 {
		return
	}
	switch path[0] {
	case "R":
		c.facing = (c.facing + 1) % 4
	case "L":
		c.facing = (c.facing + 3) % 4
	default:
		for i := lib.Int(path[0]); i > 0; i-- {
			ns, nf, v := c.Next()
			if v != '.' {
				break
			}
			if ns == c.pos {
				fmt.Println(c.pos, c.facing, "not moved")
				break
			}
			mp, mf := c.PositionFacing()
			m.Set(mp, int8(">v<^"[mf]))
			c.pos = ns
			c.facing = nf
		}
	}
	c.simulate(path[1:], m)

	mp, mf := c.PositionFacing()
	m.Set(mp, int8(">v<^"[mf]))
}

func (m *M) nextPart1() lib.Coord {
	p := m.pos
	p.Add(dir(m.facing))
	v, has := m.Get(p)
	if has && v == '#' {
		return m.pos
	}
	if !has {
		start := m.pos
		switch m.facing {
		case 0:
			start.X = m.Bounds()[0].X
		case 1:
			start.Y = m.Bounds()[0].Y
		case 2:
			start.X = m.Bounds()[1].Y
		case 3:
			start.Y = m.Bounds()[1].Y
		default:
			panic("wrong facing")
		}
		for {
			if v, has := m.Get(start); has {
				return lib.Ternary(v == '.', start, m.pos)
			}
			start.Add(dir(m.facing))
		}
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
