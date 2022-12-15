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
	ClosestBeacon     lib.Coord
	ClosestBeaconDist int
}

func (p Pos) dist(c lib.Coord) int {
	return dist(p.Coord, c)
}

func dist(a, c lib.Coord) int {
	return lib.AbsDiff(a.X, c.X) + lib.AbsDiff(a.Y, c.Y)
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

func (p Pos) coversArea(c lib.Coord, size lib.Coord) bool {
	midArea := lib.Coord{X: c.X + size.X/2, Y: c.Y + size.Y/2}
	if dist(p.Coord, midArea) < p.ClosestBeaconDist {
		return true
	}
	d := lib.Coord{X: p.X - midArea.X, Y: p.Y - midArea.Y}
	return lib.AbsDiff(d.X, size.X) < p.ClosestBeaconDist && lib.AbsDiff(d.Y, size.Y) < p.ClosestBeaconDist
}

func (p Pos) covers(c lib.Coord) bool {
	return p.dist(c) <= p.ClosestBeaconDist
}

func (p Pos) covered2(y bool, size int) (out []bool) {
	d := p.ClosestBeaconDist
	out = make([]bool, size)

	if y {
		if p.Coord.X >= 0 && p.Coord.X < size {
			out[p.Coord.X] = true
		}
		for i := 1; i < d && (p.Coord.X+i < size || p.Coord.X-i >= 0); i++ {
			if p.Coord.X+i < size {
				out[p.Coord.X+i] = true
			}
			if p.Coord.X-i >= 0 {
				out[p.Coord.X-i] = true
			}
		}
	} else {
		if p.Coord.Y >= 0 && p.Coord.Y < size {
			out[p.Coord.Y] = true
		}
		for i := 1; i < d && (p.Coord.Y+i < size || p.Coord.Y-i >= 0); i++ {
			if p.Coord.Y+i < size {
				out[p.Coord.Y+i] = true
			}
			if p.Coord.Y-i >= 0 {
				out[p.Coord.Y-i] = true
			}
		}
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
		mapp.Set(sensor, Pos{Coord: sensor, IsSensor: true, ClosestBeacon: beacon, ClosestBeaconDist: dist(sensor, beacon)})
		mapp.Set(beacon, Pos{Coord: beacon, IsBeacon: true})
	})
	fmt.Println(mapp.Draw(vis))

	// before := mapp.Len() - 1 // minus +
	// maxy := mapp.Bounds()[1].Y
	// originalHeight := mapp.Height()

	max := 20
	// max := 4000000

	squaresDimCount := 2
	squares := make([]struct {
		Sensors []Pos
		Beacons []Pos
	}, squaresDimCount*squaresDimCount)
	squareSize := lib.Coord{X: max / squaresDimCount, Y: max / squaresDimCount}

	sqnr := 0
	for x := 0; x < max; x += max / squaresDimCount {
		// fmt.Println(x)
		for y := 0; y < max; y += max / squaresDimCount {
			squareBase := lib.Coord{X: x, Y: y}
			mapp.Each(func(p Pos) bool {
				if p.IsSensor && p.coversArea(squareBase, squareSize) {
					squares[sqnr].Sensors = append(squares[sqnr].Sensors, p)
				}
				if p.IsBeacon && p.X >= x && p.Y >= y && p.X < x+max/squaresDimCount && p.Y < y+max/squaresDimCount {
					squares[sqnr].Beacons = append(squares[sqnr].Beacons, p)
				}
				return true
			})
			sqnr++
		}
	}

	sqnr = 0
	for x := 0; x < max; x += max / squaresDimCount {
		// fmt.Println(x)
		for y := 0; y < max; y += max / squaresDimCount {
			dim := max / squaresDimCount
			pixels := make([]bool, dim*dim)

			for _, s := range squares[sqnr].Sensors {
				for sx := 0; sx < dim; sx++ {
					for sy := 0; sy < dim; sy++ {
						c := s.covers(lib.Coord{X: x + sx, Y: y + sy})
						if c {
							pixels[sx*dim+sy] = true
						}
					}
				}
			}
			for _, b := range squares[sqnr].Beacons {
				sx := b.X - x
				sy := b.Y - y
				pixels[sx*dim+sy] = true
			}

			i := FalseIndex(pixels)
			if i >= 0 {
				sx := i / dim
				sy := i % dim
				fmt.Printf("uncovered x=%d y=%d\n", x+sx, y+sy)
			}
			sqnr++
		}
	}

	fmt.Println("done", squares)
}

func FalseIndex(bs []bool) int {
	for i, b := range bs {
		if !b {
			return i
		}
	}
	return -1
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
