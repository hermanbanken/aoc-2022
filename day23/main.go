package main

import (
	"aoc/lib"
	"fmt"
)

type M struct {
	lib.InfinityMap[*Elf]
}
type Elf struct {
	lib.Coord
	move     int
	proposal lib.Coord
}

func (e Elf) Check(proposal int) [3]lib.Coord {
	pos := e.Coord
	switch moves[(e.move+proposal)%len(moves)] {
	case 'N':
		return [3]lib.Coord{pos.AddR(lib.Coord{X: 0, Y: -1}), pos.AddR(lib.Coord{X: -1, Y: -1}), pos.AddR(lib.Coord{X: 1, Y: -1})}
	case 'S':
		return [3]lib.Coord{pos.AddR(lib.Coord{X: 0, Y: 1}), pos.AddR(lib.Coord{X: -1, Y: 1}), pos.AddR(lib.Coord{X: 1, Y: 1})}
	case 'W':
		return [3]lib.Coord{pos.AddR(lib.Coord{X: -1, Y: 0}), pos.AddR(lib.Coord{X: -1, Y: -1}), pos.AddR(lib.Coord{X: -1, Y: 1})}
	case 'E':
		return [3]lib.Coord{pos.AddR(lib.Coord{X: 1, Y: 0}), pos.AddR(lib.Coord{X: 1, Y: -1}), pos.AddR(lib.Coord{X: 1, Y: 1})}
	default:
		panic("invalid move")
	}
}

func (e Elf) Dir(proposal int) lib.Coord {
	pos := e.Coord
	switch moves[(e.move+proposal)%len(moves)] {
	case 'N':
		return pos.AddR(lib.Coord{X: 0, Y: -1})
	case 'S':
		return pos.AddR(lib.Coord{X: 0, Y: 1})
	case 'W':
		return pos.AddR(lib.Coord{X: -1, Y: 0})
	case 'E':
		return pos.AddR(lib.Coord{X: 1, Y: 0})
	default:
		panic("invalid move")
	}
}

const moves = "NSWE"

func main() {
	var m *M = &M{}

	var y = 0
	lib.EachLine(func(line string) {
		for x, char := range line {
			if char == '#' {
				m.Set(lib.Coord{X: x, Y: y}, &Elf{move: 0, Coord: lib.Coord{X: x, Y: y}})
			}
		}
		y += 1
	})

	anyMoved := true
	fmt.Println(m.Draw(draw) + "\n\n")
	for i := 0; i < 10 && anyMoved; i++ {
		var proposals = map[lib.Coord][]*Elf{}
		m.EachCoord(func(c lib.Coord, e *Elf) bool {
			var allFree = true
			var moveFound = false
			var proposal lib.Coord
			for i := 0; i < 4; i++ {
				pos := e.Check(i)
				anySet := m.IsSet(pos[0]) || m.IsSet(pos[1]) || m.IsSet(pos[2])
				if anySet {
					allFree = false
				} else if !moveFound {
					moveFound = true
					proposal = e.Dir(i)
				}
			}
			if !allFree && moveFound {
				e.proposal = proposal
				proposals[e.proposal] = append(proposals[e.proposal], e)
			}
			e.move += 1
			return true
		})

		for c, p := range proposals {
			if len(p) == 1 {
				m.Delete(p[0].Coord)
				m.Set(c, p[0])
				p[0].Coord = c
			}
		}
		anyMoved = len(proposals) > 0
		fmt.Println(m.Draw(draw) + "\n\n")
	}
	m.FreshBounds()
	// not 6806
	fmt.Println("part1 area", m.BoundsArea(), "count", m.Len(), "uncovered area", m.BoundsArea()-m.Len())

}

func draw(e *Elf) byte {
	if e != nil {
		return '#'
	}
	return '.'
}
