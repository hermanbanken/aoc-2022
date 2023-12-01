package main

import (
	"aoc/lib"
	"fmt"
	"strings"
)

func parseCoord(xy string) lib.Coord {
	xy = strings.NewReplacer("x=", "", "y=", "").Replace(xy)
	parts := strings.Split(xy, ", ")
	c := lib.Coord{X: lib.Int(parts[0]), Y: lib.Int(parts[1])}
	return c
}

type Pos struct {
	lib.Coord
	IsSensor          bool
	IsBeacon          bool
	IsEdge            bool
	ClosestBeacon     lib.Coord
	ClosestBeaconDist int
}

func (p Pos) dist(c lib.Coord) int {
	return dist(p.Coord, c)
}

func dist(a, c lib.Coord) int {
	return lib.AbsDiff(a.X, c.X) + lib.AbsDiff(a.Y, c.Y)
}

func (p Pos) justOutside() (out []lib.Coord) {
	d := p.ClosestBeaconDist + 1
	top := lib.Coord{X: p.X, Y: p.Y - d}
	bottom := lib.Coord{X: p.X, Y: p.Y + d}
	left := lib.Coord{X: p.X - d, Y: p.Y}
	right := lib.Coord{X: p.X + d, Y: p.Y}
	c := top
	for c != right {
		c.Add(lib.Coord{X: 1, Y: 1})
		out = append(out, c)
	}
	for c != bottom {
		c.Add(lib.Coord{X: -1, Y: 1})
		out = append(out, c)
	}
	for c != left {
		c.Add(lib.Coord{X: -1, Y: -1})
		out = append(out, c)
	}
	for c != top {
		c.Add(lib.Coord{X: 1, Y: -1})
		out = append(out, c)
	}
	return out
}

func (p Pos) covered(y int) (out []lib.Coord) {
	remainder := p.ClosestBeaconDist - lib.AbsDiff(p.Coord.Y, y) + 1
	if remainder < 0 {
		return
	}
	out = append(out, lib.Coord{Y: y, X: p.Coord.X})
	for i := 1; i < remainder; i++ {
		out = append(out, lib.Coord{Y: y, X: p.Coord.X + i})
		out = append(out, lib.Coord{Y: y, X: p.Coord.X - i})
	}
	return
}

func (p Pos) covers(c lib.Coord) bool {
	return p.dist(c) <= p.ClosestBeaconDist
}

func vis(b Pos) byte {
	if b.IsSensor {
		return 'S'
	}
	if b.IsBeacon {
		return 'B'
	}
	if b.IsEdge {
		return '#'
	}
	return '.'
}

func main() {
	var mapp = lib.InfinityMap[Pos]{}.SetDefault(Pos{})

	max := 4000000
	var outside []lib.Coord
	lib.EachLine(func(line string) {
		parts := strings.Split(line, ": ")
		sensor := parseCoord(strings.TrimPrefix(parts[0], "Sensor at "))
		beacon := parseCoord(strings.TrimPrefix(parts[1], "closest beacon is at "))
		sp := Pos{Coord: sensor, IsSensor: true, ClosestBeacon: beacon, ClosestBeaconDist: dist(sensor, beacon)}
		mapp.Set(sensor, sp)
		mapp.Set(beacon, Pos{Coord: beacon, IsBeacon: true})
		for _, outsidep := range sp.justOutside() {
			if outsidep.X > max || outsidep.X < 0 {
				continue
			}
			if outsidep.Y > max || outsidep.Y < 0 {
				continue
			}
			outside = append(outside, outsidep)
			if max <= 20 {
				edgep, _ := mapp.Get(outsidep)
				edgep.IsEdge = true
				mapp.Set(outsidep, edgep)
			}
		}
		fmt.Println(sp)
	})
	fmt.Println("mapp.Len()", mapp.Len())
	fmt.Println("len(outside)", len(outside))
	if max <= 20 {
		fmt.Println(mapp.Draw(vis))
	}

	for _, outside := range outside {
		if v, _ := mapp.Get(outside); !v.IsBeacon {
			covered := false
			mapp.Each(func(p Pos) bool {
				covered = covered || p.covers(outside)
				return !covered
			})
			if !covered {
				fmt.Println("uncovered", outside, "freq", outside.X*max+outside.Y)
				break
			}
		}
	}
}
