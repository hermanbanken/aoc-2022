package main

import (
	"aoc/lib"
	"log"
	"strconv"
	"strings"
)

type Game = []Round
type Round struct {
	R, G, B int
}

var maxRed = 12
var maxGreen = 13
var maxBlue = 14

func main() {
	sum := 0
	powerSum := 0
	games := []Game{}
	lib.EachLine(func(line string) {
		line = line[strings.Index(line, ":")+2:]
		if line == "" {
			return
		}
		sets := strings.Split(line, "; ")
		game := Game{}
		for _, set := range sets {
			colors := strings.Split(set, ", ")
			round := Round{}
			for _, color := range colors {
				parts := strings.Split(color, " ")
				digit, _ := strconv.Atoi(parts[0])
				round.Set(parts[1], digit)
			}
			game = append(game, round)
		}
		games = append(games, game)
		log.Println(game)
		minimum := Round{}
		for _, set := range game {
			if minimum.R < set.R {
				minimum.R = set.R
			}
			if minimum.G < set.G {
				minimum.G = set.G
			}
			if minimum.B < set.B {
				minimum.B = set.B
			}
		}
		powerSum += minimum.R * minimum.G * minimum.B
		for _, set := range game {
			if set.R > maxRed {
				return
			}
			if set.G > maxGreen {
				return
			}
			if set.B > maxBlue {
				return
			}
		}
		sum += len(games)
	})
	log.Println(sum)
	log.Println(powerSum)
}

func (r *Round) Set(str string, digit int) {
	switch str {
	case "red":
		r.R = digit
	case "green":
		r.G = digit
	case "blue":
		r.B = digit
	}
}
