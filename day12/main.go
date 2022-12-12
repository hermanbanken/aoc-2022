package main

import (
	"aoc/lib"
	"bufio"
	"log"
	"strings"
)

type Grid struct {
	lib.Grid
}

func (g Grid) canMoveUp(posA, posB int) bool {
	levelA := g.level(posA)
	levelB := g.level(posB)
	return levelB < levelA || lib.AbsDiff(int(levelA), int(levelB)) <= 1
}

func (g Grid) canMoveDown(posA, posB int) bool {
	levelA := g.level(posA)
	levelB := g.level(posB)
	return levelB > levelA || lib.AbsDiff(int(levelA), int(levelB)) <= 1
}

func (g Grid) level(pos int) byte {
	level := g.D[pos]
	if level == 'S' {
		level = 'a'
	}
	if level == 'E' {
		level = 'z'
	}
	return level
}

func main() {
	r := lib.Reader()
	defer r.Close()
	scanner := bufio.NewScanner(r)

	var grid Grid
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			break
		}
		grid.D = append(grid.D, line...)
		grid.W = len(line)
	}
	grid.H = len(grid.D) / grid.W

	grid.CanMove = grid.canMoveUp
	log.Println("up", grid.Dijkstra(grid.Pos('S'), func(pos int) bool { return grid.D[pos] == 'E' }))
	grid.CanMove = grid.canMoveDown
	log.Println("down", grid.Dijkstra(grid.Pos('E'), func(pos int) bool { return grid.D[pos] == 'a' || grid.D[pos] == 'S' }))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
