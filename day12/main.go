package main

import (
	"aoc/lib"
	"bufio"
	"bytes"
	"fmt"
	"log"
	"strings"
)

type Grid struct {
	d    []byte
	w, h int
}

func (g Grid) Pos(elem byte) int {
	return bytes.IndexByte(g.d, elem)
}

func (g Grid) Dijkstra(start, end byte) int {
	var dist []int = make([]int, len(g.d))
	var prev []int = make([]int, len(g.d))
	for i := range dist {
		dist[i] = -1
	}
	source := g.Pos(start)
	dist[source] = 0
	heads := []int{source}

	for {
		newHeads := []int{}
		fmt.Println("heads", heads)
		for _, head := range heads {
			for _, m := range g.Moves(head) {
				if dist[m] == -1 {
					if g.d[m] == 'E' {
						return dist[head] + 1
					}
					dist[m] = dist[head] + 1
					prev[m] = head
					newHeads = append(newHeads, m)
				}
			}
		}
		if len(newHeads) == 0 {
			g.Visited(dist, prev, heads)
			fmt.Println("dead end from", heads)
			for _, h := range heads {
				fmt.Print(string(g.d[h]))
			}
			fmt.Println()
			break
		}
		heads = newHeads
	}
	return -1
}

func (g Grid) Follow(prev []int, head int) (out []int) {
	for prev[head] != 0 {
		out = append(out, prev[head])
		head = prev[head]
	}
	return
}

func (g Grid) Visited(dist []int, prev []int, heads []int) {
	trail := []int{}
	for _, h := range heads {
		trail = append(trail, g.Follow(prev, h)...)
	}

	for y := 0; y < g.h; y++ {
		for i, v := range dist[y*g.w : (y+1)*g.w] {
			if lib.Contains(heads, y*g.w+i) {
				fmt.Print("X")
			} else if lib.Contains(trail, y*g.w+i) {
				fmt.Print("x")
			} else if v == -1 {
				fmt.Print(" ")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func (g Grid) CanMove(levelA, levelB byte) bool {
	if levelA == 'S' {
		levelA = 'a'
	}
	if levelB == 'E' {
		levelB = 'z'
	}
	return levelA > levelB || lib.AbsDiff(int(levelA), int(levelB)) <= 1
}

func (g Grid) Moves(pos int) (out []int) {
	value := g.d[pos]
	x := pos % g.w
	y := pos / g.w
	if x < g.w-1 && g.CanMove(value, g.d[pos+1]) {
		out = append(out, pos+1) // right
	}
	if x > 0 && g.CanMove(value, g.d[pos-1]) {
		out = append(out, pos-1) // left
	}
	if y > 0 && g.CanMove(value, g.d[pos-g.w]) {
		out = append(out, pos-g.w) // up
	}
	if y < g.h-1 && g.CanMove(value, g.d[pos+g.w]) {
		out = append(out, pos+g.w) // down
	}
	return
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
		grid.d = append(grid.d, line...)
		grid.w = len(line)
	}
	grid.h = len(grid.d) / grid.w

	log.Println(grid.Dijkstra('S', 'E'))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
