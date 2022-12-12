package lib

import (
	"fmt"
)

type Grid struct {
	W, H    int
	CanMove func(posA, posB int) bool
}

func (g Grid) Dijkstra(start int, end func(int) bool) (distance int, heads []int, dist []int, prev []int) {
	dist = make([]int, g.H*g.W)
	prev = make([]int, g.H*g.W)
	for i := range dist {
		dist[i] = -1
	}
	dist[start] = 0
	heads = []int{start}

	for {
		newHeads := []int{}
		for _, head := range heads {
			for _, m := range g.Moves(head) {
				if dist[m] == -1 {
					if end(m) {
						return dist[head] + 1, heads, dist, prev
					}
					dist[m] = dist[head] + 1
					prev[m] = head
					newHeads = append(newHeads, m)
				}
			}
		}
		if len(newHeads) == 0 {
			break
		}
		heads = newHeads
	}
	return -1, heads, dist, prev
}

func (g Grid) Follow(prev []int, head int) (out []int) {
	for prev[head] != 0 {
		out = append(out, prev[head])
		head = prev[head]
	}
	return
}

func (g Grid) Visualize(dist []int, prev []int, heads []int) {
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
