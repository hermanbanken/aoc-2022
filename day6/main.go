package main

import (
	"aoc/lib"
	"fmt"
	"regexp"
	"strings"
)

type Part struct {
	lib.Coord
	Length int
	Value  int
}

var nrPos = []Part{}

func main() {
	in := lib.Lines()
	timeStr := strings.TrimPrefix(in[0], "Time: ")
	distStr := strings.TrimPrefix(in[1], "Distance: ")
	times := regexp.MustCompile(`\s+`).Split(strings.TrimSpace(timeStr), -1)
	dists := regexp.MustCompile(`\s+`).Split(strings.TrimSpace(distStr), -1)

	power := 1
	output := 0
	for i, time := range times {
		dist := dists[i]
		max := lib.Int(time)

		fmt.Println("game", i+1, "t=", time, "d=", dist)
		// each release time
		for t_bat := 0; t_bat <= max; t_bat++ {
			movingTime := max - t_bat

			// speeding up time
			// speedingTime := lib.Min(movingTime, t_bat)
			// phaseA := float64(speedingTime) * float64(speedingTime) / 2
			// coasting time
			// coastingTime := max - t_bat - speedingTime + 1
			// phaseB := coastingTime * speedingTime
			calc := movingTime * t_bat
			// fmt.Printf("  t=%d, moving=%d, speeding=%d, coasting=%d, calc=%d", t_bat, movingTime, speedingTime, coastingTime, calc)
			fmt.Printf("  t=%d, moving=%d, calc=%d", t_bat, movingTime, calc)
			if calc > lib.Int(dists[i]) {
				output++
				fmt.Print(" .")
			} else {
				fmt.Print(" x")
			}
			fmt.Println()
		}
		power *= output
		output = 0
	}
	fmt.Println(power)

}
