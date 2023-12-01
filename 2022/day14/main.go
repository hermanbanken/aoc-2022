package main

import (
	"aoc/lib"
	"fmt"
	"math"
	"strings"
	"time"
)

func main() {
	var sand = lib.Coord{X: 500, Y: 0}
	var mapp = lib.InfinityMap[byte]{}.SetDefault('.')
	mapp.Set(sand, '+')

	lib.EachLine(func(line string) {
		coords := strings.Split(line, " -> ")
		prev := parse(coords[0])
		for _, coord := range coords[1:] {
			next := parse(coord)
			for prev != next {
				mapp.Set(prev, '#')
				prev.Add(next.Dir(prev))
			}
		}
		mapp.Set(prev, '#')
	})
	fmt.Println(mapp.Draw(func(b byte) byte { return b }))
	before := mapp.Len() - 1 // minus +
	maxy := mapp.Bounds()[1].Y
	originalHeight := mapp.Height()

	for {
		particle := Particle{sand}
		for particle.Down(mapp) && particle.Coord != sand && particle.Y < maxy+1 {
		}
		mapp.Set(particle.Coord, 'o')
		if particle.Coord == sand {
			fmt.Println("done", mapp.Len()-before)
			break
		}
		if mapp.Len()%int(math.Sqrt(float64(mapp.Len()))) == 0 {
			fmt.Println(mapp.Draw(func(b byte) byte { return b }) + repeat("\n", originalHeight+10-mapp.Height()))
			time.Sleep(30 * time.Millisecond)
		}
	}
	fmt.Println(mapp.Draw(func(b byte) byte { return b }) + repeat("\n", originalHeight+10-mapp.Height()))

}

func parse(xcy string) lib.Coord {
	xy := strings.Split(xcy, ",")
	return lib.Coord{X: lib.Int(xy[0]), Y: lib.Int(xy[1])}
}

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
