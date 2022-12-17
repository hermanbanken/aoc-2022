package main

import (
	"aoc/lib"
	"fmt"
	"strings"
)

type Shape struct {
	origin lib.Coord
	lib.InfinityMap[byte]
}

func (s Shape) Intersects(move lib.Coord, grid lib.InfinityMap[byte]) bool {
	s.origin.Add(move)
	var hasPixel bool
	s.EachCoord(func(c lib.Coord, b byte) bool {
		c.Add(lib.Coord{X: s.origin.X, Y: s.origin.Y})
		_, hasPixel = grid.Get(c)
		return !hasPixel
	})
	return hasPixel
}

func (s Shape) AsGrid() lib.InfinityMap[byte] {
	cop := lib.InfinityMap[byte]{}
	s.EachCoord(func(c lib.Coord, b byte) bool {
		c.Add(s.origin)
		cop.Set(c, b)
		return true
	})
	return cop
}

func (s Shape) AddTo(grid *lib.InfinityMap[byte]) {
	s.EachCoord(func(c lib.Coord, b byte) bool {
		c.Add(s.origin)
		grid.Set(c, b)
		return true
	})
}

func makeShape(lines ...string) Shape {
	s := Shape{}
	for y, l := range lines {
		for x, c := range []byte(l) {
			if c == '#' {
				s.Set(lib.Coord{X: x, Y: y}, c)
			}
		}
	}
	return s
}

func main() {
	jet := lib.Lines()[0]
	rocks := "-+L|#"
	rockShapes := map[string]Shape{
		"-": makeShape("####"),
		"+": makeShape(".#.", "###", ".#."),
		"L": makeShape("..#", "..#", "###"),
		"|": makeShape("#", "#", "#", "#"),
		"#": makeShape("##", "##"),
	}

	var grid lib.InfinityMap[byte]
	grid.SetDefault('.')
	grid.Set(lib.Coord{X: 0, Y: 0}, '-')
	grid.Set(lib.Coord{X: 1, Y: 0}, '-')
	grid.Set(lib.Coord{X: 2, Y: 0}, '-')
	grid.Set(lib.Coord{X: 3, Y: 0}, '-')
	grid.Set(lib.Coord{X: 4, Y: 0}, '-')
	grid.Set(lib.Coord{X: 5, Y: 0}, '-')
	grid.Set(lib.Coord{X: 6, Y: 0}, '-')

	fmt.Println("repeat", len(jet)*len(rocks), len(jet))
	beforeRepeatHeight := 0
	iterations := 1000000000000 / 50455
	iterationsRemainder := 1000000000000 % 50455
	iterationHeight := 0

	jetIdx := 0
	for rockIdx := 0; rockIdx < 1000000000000; rockIdx++ {
		if rockIdx > 0 && rockIdx%(50455*4) == 0 && beforeRepeatHeight > 0 {
			afterRepeatHeight := -grid.Bounds()[0].Y
			iterationHeight = (afterRepeatHeight - beforeRepeatHeight) / 4
			fmt.Println(afterRepeatHeight, beforeRepeatHeight)
			fmt.Println("iterationHeight=", iterationHeight)
		}
		if rockIdx > 0 && rockIdx%(50455*4) == 0 && beforeRepeatHeight == 0 {
			beforeRepeatHeight = -grid.Bounds()[0].Y
			fmt.Println("beforeRepeatHeight=", beforeRepeatHeight)
		}
		if rockIdx%50455 == iterationsRemainder && iterationHeight > 0 {
			fmt.Println("height=", -grid.Bounds()[0].Y+(iterations-rockIdx/50455)*iterationHeight)
			return
		}
		if rockIdx%1000 == 0 {
			progress := float64(rockIdx) * 100 / 1000000000000
			fmt.Println(progress)
		}
		t := string(rocks[rockIdx%len(rocks)])
		s := rockShapes[t]
		// fmt.Println(grid.Bounds())
		s.origin.Y = grid.Bounds()[0].Y - 4 - s.Bounds()[1].Y
		s.origin.X = 2

		// fmt.Println("Start", t)
		// fmt.Println(visualize(grid, s.AsGrid()))
		for {
			// move by jet
			b := s.Bounds()
			b[0].Add(s.origin)
			b[1].Add(s.origin)
			isRight := jet[jetIdx%len(jet)] == '>'
			jetIdx++
			if isRight && b[1].X < 6 && !s.Intersects(lib.Coord{X: 1, Y: 0}, grid) {
				s.origin.X += 1
				// fmt.Println("right")
			} else if !isRight && b[0].X > 0 && !s.Intersects(lib.Coord{X: -1, Y: 0}, grid) {
				s.origin.X -= 1
				// fmt.Println("left")
			} else {
				// fmt.Println("blocked")
			}

			// fall
			if !s.Intersects(lib.Coord{X: 0, Y: 1}, grid) {
				s.origin.Y += 1
				// fmt.Println("down")
			} else {
				// fmt.Println("done")
				s.AddTo(&grid)
				break
			}
			// // done after fall?
			// if s.Intersects(lib.Coord{X: 0, Y: 1}, grid) {
			// 	fmt.Println("done")
			// 	s.AddTo(grid)
			// 	break
			// }
		}
		// fmt.Println(visualize(grid, s.AsGrid()))
	}
	// fmt.Println(visualize(grid, lib.InfinityMap[byte]{}))
	fmt.Println(grid.Bounds())
	_ = jet
	_ = rocks
	_ = rockShapes
}

func visualize(grid ...lib.InfinityMap[byte]) string {
	b := grid[0].Bounds()

	var sb = strings.Builder{}
	for y := b[0].Y - 5; y < b[1].Y+1; y++ {
		for x := b[0].X; x < b[1].X+1; x++ {
			var v byte
			var has bool
			for _, g := range grid {
				v, has = g.Get(lib.Coord{X: x, Y: y})
				if has {
					break
				}
			}
			if !has {
				sb.WriteString(".")
			} else {
				sb.WriteString(string(v))
			}
		}
		sb.WriteRune('\r')
		sb.WriteRune('\n')
	}
	fmt.Print("\n")
	return sb.String()
}
