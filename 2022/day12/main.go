package main

import (
	"aoc/lib"
	"bufio"
	"bytes"
	"log"
	"strings"
)

type Grid struct {
	lib.Grid
	D []byte
}

func (g Grid) canMoveUp(posA, posB int) bool {
	levelA := g.D[posA]
	levelB := g.D[posB]
	return levelB < levelA || lib.AbsDiff(int(levelA), int(levelB)) <= 1
}

func (g Grid) canMoveDown(posA, posB int) bool {
	levelA := g.D[posA]
	levelB := g.D[posB]
	return levelB > levelA || lib.AbsDiff(int(levelA), int(levelB)) <= 1
}

func main() {
	r := lib.Reader()
	defer r.Close()
	scanner := bufio.NewScanner(r)

	var grid Grid
	var start, end int
	for scanner.Scan() {
		if line := scanner.Text(); strings.TrimSpace(line) != "" {
			grid.D = append(grid.D, line...)
			grid.W = len(line)
		}
	}
	grid.H = len(grid.D) / grid.W

	start = bytes.IndexByte(grid.D, 'S')
	end = bytes.IndexByte(grid.D, 'E')
	grid.D[start] = 'a'
	grid.D[end] = 'z'

	grid.CanMove = grid.canMoveUp
	d, heads, dist, prev := grid.Dijkstra(start, func(pos int) bool { return pos == end })
	grid.Visualize(dist, prev, heads)
	log.Println("up", d)

	grid.CanMove = grid.canMoveDown
	d, heads, dist, prev = grid.Dijkstra(end, func(pos int) bool { return grid.D[pos] == 'a' })
	grid.Visualize(dist, prev, heads)
	log.Println("down", d)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
