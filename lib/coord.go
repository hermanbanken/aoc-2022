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

func (p Coord) AddR(o Coord) Coord {
	return Coord{X: p.X + o.X, Y: p.Y + o.Y}
}

func (p *Coord) Mult(i int) {
	p.X *= i
	p.Y *= i
}
func (p Coord) MultR(i int) Coord {
	return Coord{X: p.X * i, Y: p.Y * i}
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
	return strconv.Itoa(p.X) + "," + strconv.Itoa(p.Y)
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
		m.include(c)
	}
	m.data[c] = val
}

func (m *InfinityMap[T]) include(c Coord) {
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

func (m *InfinityMap[T]) FreshBounds() [2]Coord {
	var first = true
	for c, _ := range m.data {
		if first {
			first = false
			m.bounds[0] = c
			m.bounds[1] = c
		}
		m.include(c)
	}
	return m.bounds
}

func (m *InfinityMap[T]) Delete(c Coord) {
	delete(m.data, c)
}

func (m *InfinityMap[T]) Each(fn func(T) bool) {
	for _, c := range m.data {
		cont := fn(c)
		if !cont {
			break
		}
	}
}

func (m *InfinityMap[T]) EachCoord(fn func(Coord, T) bool) {
	for k, c := range m.data {
		cont := fn(k, c)
		if !cont {
			break
		}
	}
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

func (m InfinityMap[T]) Width() int {
	return m.bounds[1].X - m.bounds[0].X
}
func (m InfinityMap[T]) Height() int {
	return m.bounds[1].Y - m.bounds[0].Y
}
func (m *InfinityMap[T]) BoundsArea() int {
	w := m.bounds[1].X - m.bounds[0].X
	h := m.bounds[1].Y - m.bounds[0].Y
	return w * h
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

func (m InfinityMap[T]) IsSet(c Coord) bool {
	_, has := m.data[c]
	return has
}
