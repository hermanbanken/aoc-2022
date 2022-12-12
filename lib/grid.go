package lib

import (
	"bytes"
	"fmt"
)

type Grid struct {
	D       []byte
	W, H    int
	CanMove func(posA, posB int) bool
}

func (g Grid) Pos(elem byte) int {
	return bytes.IndexByte(g.D, elem)
}

func (g Grid) Dijkstra(start int, end func(int) bool) int {
	var dist []int = make([]int, len(g.D))
	var prev []int = make([]int, len(g.D))
	for i := range dist {
		dist[i] = -1
	}
	dist[start] = 0
	heads := []int{start}

	for {
		newHeads := []int{}
		for _, head := range heads {
			for _, m := range g.Moves(head) {
				if dist[m] == -1 {
					if end(m) {
						g.Visited(dist, prev, heads)
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
				fmt.Print(string(g.D[h]))
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

	for y := 0; y < g.H; y++ {
		for i, v := range dist[y*g.W : (y+1)*g.W] {
			if Contains(heads, y*g.W+i) {
				fmt.Print("X")
			} else if Contains(trail, y*g.W+i) {
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

func (g Grid) Moves(pos int) (out []int) {
	x := pos % g.W
	y := pos / g.W
	if x < g.W-1 && g.CanMove(pos, pos+1) {
		out = append(out, pos+1) // right
	}
	if x > 0 && g.CanMove(pos, pos-1) {
		out = append(out, pos-1) // left
	}
	if y > 0 && g.CanMove(pos, pos-g.W) {
		out = append(out, pos-g.W) // up
	}
	if y < g.H-1 && g.CanMove(pos, pos+g.W) {
		out = append(out, pos+g.W) // down
	}
	return
}
