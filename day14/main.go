package main

import (
	"aoc/lib"
	"fmt"
	"sort"
	"strings"
)

type Tile [2]int

func weight(c int) int {
	switch c {
	case int('#'):
		return 10
	case int('O'):
		return 20
	case int('.'):
		return 0
	}
	panic(fmt.Sprintf("invalid %v", rune(c)))
}

func (tile Tile) Less(other Tile) bool {
	return Less(tile, other)
}

func Less(tile Tile, other Tile) bool {
	if tile[1] == other[1] {
		return weight(tile[0]) > weight(other[0])
	}
	return tile[1] < other[1]
}

func weighM(m Matrix[Tile]) (sum int) {
	height := len(m.data) / m.w
	m.ForEach(func(x, y int, t Tile) {
		if t[0] == 'O' {
			sum += height - y
		}
	})
	return sum
}
func print(t Tile) byte { return byte(t[0]) }

func main() {
	lines := lib.Lines()
	m := Matrix[Tile]{
		data: lib.Map([]byte(strings.Join(lines, "")), func(b byte) Tile { return Tile{int(b), 0} }),
		w:    len(lines[0]),
	}
	var index = map[string]int{}
	for cycle := 0; cycle < 1_000_000_000; cycle++ {
		if cycle%1_000 == 0 {
			fmt.Println(cycle)
		}
		// fmt.Println(m.String(print))
		// fmt.Println()
		for i := 0; i < 4; i++ {
			tiltM(m)
			if cycle == 0 && i == 0 {
				fmt.Println("part1m", weighM(m), "\n ")
			}
			m.RotateCW()
		}
		fmt.Println(weighM(m))
		if first, hasIndex := index[m.String(print)]; hasIndex && (1_000_000_000-1-cycle)%(cycle-first) == 0 {
			fmt.Println(first, "repeats at", cycle)
			fmt.Println("part2", weighM(m))
			return
		}
		index[m.String(print)] = cycle

	}
}

func (m Matrix[T]) String(fn func(T) byte) string {
	str := strings.Builder{}
	for i := 0; i < len(m.data); i++ {
		str.WriteByte(fn(m.data[i]))
		if i%m.w == m.w-1 {
			str.WriteRune('\n')
		}
	}
	return str.String()
}

type Matrix[T any] struct {
	data   []T
	buffer []T
	w      int
}

func (m *Matrix[T]) RotateCW() {
	// x=2,y=0 => x=1,y=2
	// x=2,y=1 => x=0,y=2
	bw := len(m.data) / m.w
	if m.buffer == nil {
		m.buffer = make([]T, len(m.data))
	}
	for i := 0; i < len(m.data); i++ {
		x := i % m.w
		y := i / m.w
		by := x
		bx := bw - 1 - y
		m.buffer[bx+by*bw] = m.data[i]
	}
	m.w = bw
	copy(m.data, m.buffer)
}

func (m Matrix[T]) ForEach(fn func(x, y int, t T)) {
	for i := 0; i < len(m.data); i++ {
		fn(i%m.w, i/m.w, m.data[i])
	}
}

func (m Matrix[T]) Get(x, y int) T {
	return m.data[x+y*m.w]
}
func (m Matrix[T]) Set(x, y int, t T) {
	m.data[x+y*m.w] = t
}

func (m Matrix[T]) SortColumn(i int, lessFn func(a, b T) bool) (out sort.Interface) {
	return col[T]{m, i, lessFn}
}

type col[T any] struct {
	m      Matrix[T]
	col    int
	lessFn func(a, b T) bool
}

// Len implements sort.Interface.
func (col col[T]) Len() int {
	return len(col.m.data) / col.m.w
}

// Less implements sort.Interface.
func (col col[T]) Less(i int, j int) bool {
	return col.lessFn(col.m.Get(col.col, i), col.m.Get(col.col, j))
}

// Swap implements sort.Interface.
func (col col[T]) Swap(i int, j int) {
	tmp := col.m.Get(col.col, i)
	col.m.Set(col.col, i, col.m.Get(col.col, j))
	col.m.Set(col.col, j, tmp)
}

func tiltM(m Matrix[Tile]) {
	// rank stacks
	m.ForEach(func(x, y int, t Tile) {
		if y > 0 {
			above := m.Get(x, y-1)
			m.Set(x, y, Tile{t[0], above[1] + lib.Ternary(t[0] == '#' || above[0] == '#', 1, 0)})
		} else {
			m.Set(x, y, Tile{t[0], 0})
		}
	})
	// tilt north
	for i := 0; i < m.w; i++ {
		sort.Stable(m.SortColumn(i, Less))
	}
}
