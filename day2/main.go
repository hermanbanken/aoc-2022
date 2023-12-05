package main

import (
	"aoc/lib"
	"log"
	"strconv"
	"strings"
)

type Game = []map[string]int

var maxRed = 12
var maxGreen = 13
var maxBlue = 14

func main() {
	sum := 0
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
			for _, color := range colors {
				parts := strings.Split(color, " ")
				digit, _ := strconv.Atoi(parts[0])
				game = append(game, map[string]int{parts[1]: digit})
			}
		}
		games = append(games, game)
		log.Println(game)
		for _, set := range game {
			for color, digit := range set {
				switch color {
				case "red":
					if digit > maxRed {
						return
					}
				case "blue":
					if digit > maxBlue {
						return
					}
				case "green":
					if digit > maxGreen {
						return
					}
				}
			}
		}
		sum += len(games)
	})
	log.Println(sum)
}
