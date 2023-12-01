package main

import (
	"aoc/lib"
	"fmt"
	"log"
	"sort"
	"strings"
)

type Coord struct {
	x, y, z int
}
type Cube struct {
	Coord
	isLava     bool
	isWater    bool
	sidesWater int
	sidesAir   int
	sidesLava  int
}

func (c Coord) add(o Coord) Coord {
	return Coord{c.x + o.x, c.y + o.y, c.z + o.z}
}
func (c Coord) times(i int) Coord {
	return Coord{c.x * i, c.y * i, c.z * i}
}

func main() {
	var ps = map[Coord]*Cube{}
	var psl = []*Cube{}
	lib.EachLine(func(line string) {
		xyz := strings.Split(line, ",")
		p := Cube{Coord{lib.Int(xyz[0]), lib.Int(xyz[1]), lib.Int(xyz[2])}, true, false, 0, 0, 0}
		ps[p.Coord] = &p
		psl = append(psl, &p)
	})
	originalPoints := len(ps)

	var sum = 0
	delta := []Coord{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}, {-1, 0, 0}, {0, -1, 0}, {0, 0, -1}}
	for _, p := range psl {
		if !p.isLava {
			continue
		}
		for _, d := range delta {
			v, filled := ps[p.add(d)]
			if !filled || !v.isLava {
				sum += 1
			}
			for i := 0; i < originalPoints; i++ {
				sibling := p.add(d.times(i))
				if _, filled := ps[sibling]; !filled {
					ps[sibling] = &Cube{sibling, false, true, 0, 0, 0}
				}
			}
		}
	}
	log.Println("surface area", sum)

	// sort.Slice(psl, func(i, j int) bool {
	// 	if psl[i].x < psl[j].x {
	// 		return true
	// 	}
	// 	if psl[i].y < psl[j].y {
	// 		return true
	// 	}
	// 	if psl[i].z < psl[j].z {
	// 		return true
	// 	}
	// 	return false
	// })
	// sources := []Coord{psl[0].Coord.add(Coord{-1, -1, -1})}
	// for {
	// 	sources
	// }

	for _, p := range ps {
		for _, d := range delta {
			v, filled := ps[p.add(d)]
			if !filled {
				p.sidesAir++
			} else if v.isLava {
				p.sidesLava++
			} else if v.isWater {
				p.sidesWater++
			} else {
				panic("Huh")
			}
		}
	}

	var ms []*Cube
	for _, p := range ps {
		ms = append(ms, p)
	}

	remove := func(p *Cube) {
		for _, d := range delta {
			if v, filled := ps[p.add(d)]; filled {
				if p.isLava {
					panic("should not remove this")
				} else {
					v.sidesWater--
					v.sidesAir++
				}
			}
		}
	}

	dosort := func(length int) {
		sort.Slice(ms[0:length-1], func(i, j int) bool {
			if ms[i].isLava != ms[j].isLava {
				return !ms[i].isLava
			}
			return ms[i].sidesAir > ms[j].sidesAir
		})
	}
	dosort(len(ms))

	aftersort := func() *Cube {
		dosort(len(ms))
		return ms[0]
	}

	fmt.Println("iterating", len(ms))
	it := 0
	for (ms[0].isWater && ms[0].sidesAir > 0) || (aftersort().isWater && ms[0].sidesAir > 0) {
		fmt.Println(*ms[0])

		it++
		if it%1000 == 0 {
			fmt.Println("iteration", it, len(ms))
		}
		remove(ms[0])
		delete(ps, ms[0].Coord)
		ms = ms[1:]
		dosort(10)
	}
	for _, m := range ms {
		fmt.Println("remaining", *m)
	}

	fmt.Println("iterations", it, ms[0])
	log.Println("originalPoints", originalPoints, "points", len(ps))

	hole := map[Coord]*Cube{}
	var ext = 0
	for _, p := range ps {
		if !p.isLava {
			continue
		}
		for _, d := range delta {
			v, filled := ps[p.add(d)]
			if filled && !v.isLava {
				hole[v.Coord] = v
			}
			if !filled {
				ext += 1
			}
		}
	}

	log.Println("holes", len(hole), hole)
	// not 3246
	// not 3221
	log.Println("exterior area", ext)
}
