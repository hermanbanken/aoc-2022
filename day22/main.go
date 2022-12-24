package main

import (
	"aoc/lib"
	"fmt"
	"log"
	"math"
	"strings"
)

type M struct {
	lib.InfinityMap[int8]
	side   byte
	facing int
	pos    lib.Coord
	Dim    int
}

type Stance struct {
	lib.Coord
	facing int
}

type Point struct {
	X, Y, Z int
}

func (m M) OverEdge() (bool, lib.Coord) {
	mod := lib.Coord{X: m.pos.X % m.Dim, Y: m.pos.Y % m.Dim}
	switch m.facing % 4 {
	case 0: // right
		return mod.X == m.Dim-1, m.pos.AddR(lib.Coord{1, 0})
	case 1: // down
		return mod.Y == m.Dim-1, m.pos.AddR(lib.Coord{0, 1})
	case 2: // left
		return mod.X == 0, m.pos.AddR(lib.Coord{-1, 0})
	case 3: // up
		return mod.Y == 0, m.pos.AddR(lib.Coord{0, -1})
	default:
		panic(fmt.Sprintf("unknown facing %d", m.facing))
	}
}

func (m M) nextPart2() Stance {
	overEdge, next := m.OverEdge()
	_, exists := m.Get(next)
	// regular move
	if !overEdge || exists {
		// log.Println("regular")
		return Stance{next, m.facing}
	}
	// figure out where to go on cube space
	for _, p := range append(ortOffset(dir(m.facing), false, 1, 3), ortOffset(dir(m.facing).MultR(-3), true, 1, 3)...) {
		np := m.pos.AddR(p.MultR(m.Dim))
		if m.IsSet(np) {
			// log.Println("ort", dir((p.facing+4)%4))
			return Stance{flip(np, p.facing, m.facing, m.Dim), (p.facing + m.facing + 4) % 4}
		}
	}
	for _, p := range horse(m.facing) {
		np := m.pos.AddR(p.MultR(m.Dim))
		if m.IsSet(np) {
			// log.Println("horse")
			return Stance{pivot(np, m.facing, m.Dim), (m.facing + 2) % 4}
		}
	}

	// Wrap around
	np := m.pos.AddR(dir(m.facing).MultR(-4*m.Dim + 1))
	if m.IsSet(np) {
		return Stance{np, m.facing}
	}

	// // top to bottom in input1.txt
	// for i, p := range ort(dir(m.facing).MultR(-3)) {
	// 	np := m.pos.AddR(p.MultR(m.Dim))
	// 	if m.IsSet(np) {
	// 		return Stance{flip(np, p.facing, m.facing, m.Dim), (m.facing + lib.Ternary(i == 0, 1, -1) + 4) % 4}
	// 	}
	// }
	// for i, p := range ortOffset(dir(m.facing).MultR(-1), 3) {
	// 	np := m.pos.AddR(p.MultR(m.Dim))
	// 	if m.IsSet(np) {
	// 		return Stance{flip(np, p.facing, m.facing, m.Dim), (m.facing + lib.Ternary(i == 0, 1, -1) + 4) % 4}
	// 	}
	// }
	// // top right 3 to bottom 3 in input1.txt
	// for _, p := range ortOffset(dir(m.facing).MultR(-3), 2) {
	// 	np := m.pos.AddR(p.MultR(m.Dim))
	// 	if m.IsSet(np) {
	// 		offset := np.AddR(dir(m.facing))
	// 		offset.X %= m.Dim
	// 		offset.Y %= m.Dim
	// 		base := lib.Coord{X: np.X - np.X%m.Dim, Y: np.Y - np.Y%m.Dim}
	// 		return Stance{base.AddR(offset), m.facing}
	// 	}
	// }
	// for _, p := range ortOffset(dir(m.facing).MultR(-2), 3) {
	// 	np := m.pos.AddR(p.MultR(m.Dim))
	// 	if m.IsSet(np) {
	// 		offset := np.AddR(dir(m.facing))
	// 		offset.X %= m.Dim
	// 		offset.Y %= m.Dim
	// 		base := lib.Coord{X: np.X - np.X%m.Dim, Y: np.Y - np.Y%m.Dim}
	// 		return Stance{base.AddR(offset), m.facing}
	// 	}
	// }

	panic(fmt.Sprintf("no destination available from %s facing %d\n", m.pos, m.facing))
}

// Flip moves around this axis:
// A:
//
//	x  /
//	  /
//	 /
//	/   x
//
// B:
//
//	\  x
//	 \
//	  \
//	x  \
func flip(p lib.Coord, moveFacing, originalFacing int, dim int) (o lib.Coord) {
	base := lib.Coord{Y: p.Y - p.Y%dim, X: p.X - p.X%dim}
	max := lib.Coord{Y: p.Y - p.Y%dim + dim - 1, X: p.X - p.X%dim + dim - 1}
	offset := lib.Coord{X: p.X % dim, Y: p.Y % dim}
	switch originalFacing {
	case 0: // right
		if moveFacing == -1 { // ccw
			return lib.Coord{Y: max.Y, X: base.X + offset.Y}
		} else { // cw
			return lib.Coord{Y: base.Y, X: max.X - offset.Y}
		}
	case 1: // down
		if moveFacing == -1 { // ccw
			return lib.Coord{Y: max.Y - offset.X, X: base.X}
		} else { // cw
			return lib.Coord{Y: base.Y + offset.X, X: max.X}
		}
	case 2: // left
		if moveFacing == -1 { // ccw
			return lib.Coord{X: base.X + offset.Y, Y: base.Y}
		} else { // cw
			return lib.Coord{X: max.X - offset.Y, Y: max.Y}
		}
	case 3: // up
		if moveFacing == -1 { // ccw
			return lib.Coord{Y: max.Y - offset.X, X: max.X}
		} else { // cw
			return lib.Coord{Y: base.Y + offset.X, X: base.X}
		}
	default:
		panic("unknown facing")
	}
}

func pivot(p lib.Coord, facing int, dim int) (o lib.Coord) {
	o.X = p.X
	o.Y = p.Y
	if facing == 1 || facing == 3 {
		o.X = p.X/dim*dim + dim - (p.X % dim) - 1
	} else {
		o.Y = p.Y/dim*dim + dim - (p.Y % dim) - 1
	}
	return
}

// horse moves:
// xx1x1xxx
// x1xxx1xx
// xxx0xxxx
// x1xxx1xx
// xx1x1xxx
// PORTAL RIGHT/LEFT => 2 down/up    1 left/right
// PORTAL UP/DOWN    => 2 left/right 1 up/down
func horse(facing int) []lib.Coord {
	switch facing % 4 {
	case 0, 2:
		return []lib.Coord{
			{X: 1, Y: -2}, {X: 1, Y: 2},
			{X: -1, Y: 2}, {X: -1, Y: -2},
		}
	case 1, 3:
		return []lib.Coord{
			{X: 2, Y: -1}, {X: 2, Y: 1},
			{X: -2, Y: 1}, {X: -2, Y: -1},
		}
	default:
		panic("invalid facing")
	}
}

func read(d3 bool) (m *M, steps []string) {
	m = &M{}
	m.SetDefault(' ')
	var instructions bool
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
				p := lib.Coord{X: x, Y: y}
				// pivot input1 to look like input0
				// if d3 && y > 100 {
				// 	p = rotate(p, lib.Coord{X: 50, Y: 100}, math.Pi/2)
				// }
				if d3 && y >= 150 {
					p = rotate(p, lib.Coord{X: 49, Y: 150}, -math.Pi/2).AddR(lib.Coord{1, 0})
				}
				if d3 && y >= 50 {
					p.Y -= 200
				}
				if d3 {
					p.Y += 150
				}
				m.Set(p, int8(c))
			}
			y += 1
		}
	})
	fmt.Println("drawing:", m.Bounds())
	fmt.Println(m.Draw(func(i int8) byte { return byte(i) }))
	return m, steps
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
	m, steps := read(false)
	fmt.Println(m.Dim, steps)
	// fmt.Println(m.Draw(func(b int8) byte { return byte(b) }))

	fmt.Println(m.pos)
	m.simulate(steps, true)
	part1 := m.pos.AddR(lib.Coord{X: 1, Y: 1})
	fmt.Println("part1", part1, 1000*(part1.Y)+4*(part1.X)+m.facing)

	m, steps = read(true)
	m.simulate(steps, false)
	part2 := m.pos.AddR(lib.Coord{X: 1, Y: 1})
	part2.Y -= 150
	if part2.Y < 0 {
		part2.Y += 200
	}
	// if d3 && y >= 50 {
	// 	p.Y -= 200
	// }
	// if d3 {
	// 	p.Y += 150
	// }

	// not 106189
	// not 136374
	// not 31568
	fmt.Println("part2", part2, 1000*(part2.Y)+4*(part2.X)+m.facing)
}

func (m *M) simulate(path []string, part1 bool) {
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
			if part1 {
				n = Stance{m.nextPart1(), m.facing}
			} else {
				stance := m.nextPart2()
				if v, _ := m.Get(stance.Coord); v == '.' {
					n = stance
				}
			}
			if n.Coord == m.pos {
				// fmt.Println(m.pos, m.facing, "not moved")
				break
			}
			m.pos = n.Coord
			m.facing = n.facing
			// fmt.Println(m.pos, m.facing)
		}
	}
	m.simulate(path[1:], part1)
}

func (m *M) nextPart1() lib.Coord {
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

// for every direction (sqrt((X+Y)*(X+Y)) = 1) there are 2 positions around it
func ort(dir lib.Coord) []Stance {
	return ortOffset(dir, false, 1)
}

// for every direction (sqrt((X+Y)*(X+Y)) = 1) there are 2 positions around it
func ortOffset(dir lib.Coord, flipped bool, o ...int) (out []Stance) {
	if dir.X == 0 {
		cond := dir.Y < 0
		if flipped {
			cond = !cond
		}
		for _, o := range o {
			out = append(out,
				Stance{lib.Coord{X: -o, Y: dir.Y}, lib.Ternary(cond, -1, 1)},
				Stance{lib.Coord{X: o, Y: dir.Y}, lib.Ternary(cond, 1, -1)},
			)
		}
		return out
	}
	cond := dir.X > 0
	if flipped {
		cond = !cond
	}
	for _, o := range o {
		out = append(out,
			Stance{lib.Coord{X: dir.X, Y: -o}, lib.Ternary(cond, -1, 1)},
			Stance{lib.Coord{X: dir.X, Y: o}, lib.Ternary(cond, 1, -1)},
		)
	}
	return out
}
