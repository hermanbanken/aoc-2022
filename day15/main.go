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
	fmt.Println(xy, c)
	return c
}

type Pos struct {
	lib.Coord
	IsSensor      bool
	IsBeacon      bool
	ClosestBeacon lib.Coord
}

func (p Pos) dist(c lib.Coord) int {
	return lib.AbsDiff(p.Coord.X, c.X) + lib.AbsDiff(p.Coord.Y, c.Y)
}

func (p Pos) covered(y int) (out []lib.Coord) {
	d := p.dist(p.ClosestBeacon)
	remainder := d - lib.AbsDiff(p.Coord.Y, y) + 1
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

func vis(b Pos) byte {
	if b.IsSensor {
		return 'S'
	}
	if b.IsBeacon {
		return 'B'
	}
	return '.'
}

func main() {
	var mapp = lib.InfinityMap[Pos]{}.SetDefault(Pos{})

	lib.EachLine(func(line string) {
		parts := strings.Split(line, ": ")
		sensor := parseCoord(strings.TrimPrefix(parts[0], "Sensor at "))
		beacon := parseCoord(strings.TrimPrefix(parts[1], "closest beacon is at "))
		mapp.Set(sensor, Pos{Coord: sensor, IsSensor: true, ClosestBeacon: beacon})
		mapp.Set(beacon, Pos{Coord: beacon, IsBeacon: true})
	})
	// fmt.Println(mapp.Draw(vis))

	// before := mapp.Len() - 1 // minus +
	// maxy := mapp.Bounds()[1].Y
	// originalHeight := mapp.Height()

	y := 2000000
	covered := map[lib.Coord]bool{}
	mapp.Each(func(p Pos) {
		fmt.Println(p.Coord)
		if p.IsSensor {
			pcov := p.covered(y)
			fmt.Println(p.Coord, len(pcov))
			for _, c := range pcov {
				if v, _ := mapp.Get(c); !v.IsBeacon {
					covered[c] = true
				}
			}
		}
	})
	fmt.Println(len(covered))
}

// ####B######################
type Particle struct {
	lib.Coord
}

func (p *Particle) Down(mapp lib.InfinityMap[byte]) bool {
	var options = []lib.Coord{
		{X: p.X, Y: p.Y + 1},
		{X: p.X - 1, Y: p.Y + 1},
		{X: p.X + 1, Y: p.Y + 1},
	}
options:
	for len(options) > 0 {
		if v, _ := mapp.Get(options[0]); v != '.' {
			options = options[1:]
			continue options
		}
		p.X = options[0].X
		p.Y = options[0].Y
		return true
	}
	return false
}

func repeat(str string, c int) string {
	if c < 0 {
		return ""
	}
	return strings.Repeat(str, c)
}
