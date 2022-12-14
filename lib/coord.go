package lib

import (
	"log"
	"strconv"
	"strings"
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

type InfinityMap[T any] struct {
	data     map[Coord]T
	defaultV T
	bounds   [2]Coord
}

func (m InfinityMap[T]) SetDefault(v T) InfinityMap[T] {
	m.defaultV = v
	return m
}

func (m *InfinityMap[T]) Set(c Coord, val T) {
	if m.data == nil {
		m.data = map[Coord]T{}
	}
	if len(m.data) == 0 {
		m.bounds = [2]Coord{c, c}
	} else {
		if c.X < m.bounds[0].X {
			m.bounds[0].X = c.X
		}
		if c.Y < m.bounds[0].Y {
			m.bounds[0].Y = c.Y
		}
		if c.X > m.bounds[1].X {
			m.bounds[1].X = c.X
		}
		if c.Y > m.bounds[1].Y {
			m.bounds[1].Y = c.Y
		}
	}
	m.data[c] = val
}

func (m InfinityMap[T]) Bounds() [2]Coord {
	return m.bounds
}
func (m InfinityMap[T]) Len() int {
	return len(m.data)
}

func (m InfinityMap[T]) Get(c Coord) (T, bool) {
	v, has := m.data[c]
	if !has {
		return m.defaultV, false
	}
	return v, true
}

func (m InfinityMap[T]) Height() int {
	return m.bounds[1].Y - m.bounds[0].Y
}

func (m InfinityMap[T]) Draw(fn func(T) byte) string {
	sb := strings.Builder{}
	for y := m.bounds[0].Y; y <= m.bounds[1].Y; y++ {
		for x := m.bounds[0].X; x <= m.bounds[1].X; x++ {
			b, _ := m.Get(Coord{X: x, Y: y})
			sb.WriteByte(fn(b))
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}
