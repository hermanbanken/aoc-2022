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
	proposal lib.Coord
}

func (e Elf) Check(proposal int) [3]lib.Coord {
	pos := e.Coord
	switch moves[(proposal)%len(moves)] {
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
	switch moves[(proposal)%len(moves)] {
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
				m.Set(lib.Coord{X: x, Y: y}, &Elf{Coord: lib.Coord{X: x, Y: y}})
			}
		}
		y += 1
	})

	anyMoved := true
	moveOffset := -1
	var areaRound10 = 0
	fmt.Println(m.Draw(draw) + "\n\n")
	round := 0
	for ; anyMoved; round++ {
		moveOffset += 1
		var proposals = map[lib.Coord][]*Elf{}
		m.EachCoord(func(c lib.Coord, e *Elf) bool {
			var allFree = true
			var moveFound = false
			var proposal lib.Coord
			for i := moveOffset; i < moveOffset+4; i++ {
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
		if round == 10-1 {
			m.FreshBounds()
			areaRound10 = (m.Width()+1)*(m.Height()+1) - m.Len()
			fmt.Println(m.Draw(draw) + "\n\n")
		}
	}
	// not 6806, but 4288
	fmt.Println("part1 area", areaRound10)

	fmt.Println("part2 rounds", round)

}

func draw(e *Elf) byte {
	if e != nil {
		return '#'
	}
	return '.'
}
