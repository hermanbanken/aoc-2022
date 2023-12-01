package main

import (
	"aoc/lib"
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	r := lib.Reader()
	defer r.Close()
	scanner := bufio.NewScanner(r)

	// Read stacks
	var mp [][]int
	var hidden [][]bool
	for scanner.Scan() {
		t := scanner.Text()
		if strings.TrimSpace(t) == "" {
			break
		}
		var row = []int{}
		for _, d := range strings.Split(t, "") {
			digit, _ := strconv.Atoi(d)
			row = append(row, digit)
		}
		mp = append(mp, row)
		hidden = append(hidden, make([]bool, len(row)))
	}

	g := Grid{mp, hidden}
	highest := 0
	visible := 0
	for x := 0; x < len(mp[0]); x++ {
		for y := 0; y < len(mp); y++ {
			if g.Visible(x, y) {
				visible += 1
			}
			s := g.Score(x, y)
			if s > highest {
				highest = s
			}
		}
	}
	fmt.Println(visible, highest)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

type Grid struct {
	mp     [][]int
	hidden [][]bool
}

func (g Grid) Visible(x, y int) bool {
	return g.LeftOk(x, y) || g.RightOk(x, y) || g.TopOk(x, y) || g.BottomOk(x, y)
}

func (g Grid) LeftOk(x, y int) bool {
	for px := x - 1; px >= 0; px-- {
		if g.mp[y][px] >= g.mp[y][x] {
			return false
		}
	}
	return true
}

func (g Grid) RightOk(x, y int) bool {
	for px := x + 1; px < len(g.mp[0]); px++ {
		if g.mp[y][px] >= g.mp[y][x] {
			return false
		}
	}
	return true
}

func (g Grid) TopOk(x, y int) bool {
	for py := y - 1; py >= 0; py-- {
		if g.mp[py][x] >= g.mp[y][x] {
			return false
		}
	}
	return true
}
func (g Grid) BottomOk(x, y int) bool {
	for py := y + 1; py < len(g.mp); py++ {
		if g.mp[py][x] >= g.mp[y][x] {
			return false
		}
	}
	return true
}

func (g Grid) Score(x, y int) int {
	return g.LeftScore(x, y) * g.RightScore(x, y) * g.TopScore(x, y) * g.BottomScore(x, y)
}

func (g Grid) LeftScore(x, y int) int {
	count := 0
	for px := x - 1; px >= 0; px-- {
		count += 1
		if g.mp[y][px] >= g.mp[y][x] {
			return count
		}
	}
	return count
}

func (g Grid) RightScore(x, y int) int {
	count := 0
	for px := x + 1; px < len(g.mp[0]); px++ {
		count += 1
		if g.mp[y][px] >= g.mp[y][x] {
			return count
		}
	}
	return count
}

func (g Grid) TopScore(x, y int) int {
	count := 0
	for py := y - 1; py >= 0; py-- {
		count += 1
		if g.mp[py][x] >= g.mp[y][x] {
			return count
		}
	}
	return count
}
func (g Grid) BottomScore(x, y int) int {
	count := 0
	for py := y + 1; py < len(g.mp); py++ {
		count += 1
		if g.mp[py][x] >= g.mp[y][x] {
			return count
		}
	}
	return count
}

func (g Grid) Subset() Grid {
	if len(g.mp)-2 == 0 {
		return Grid{}
	}
	var newmp = make([][]int, len(g.mp)-2)
	for i, row := range g.mp {
		if i == 0 || i >= len(newmp) {
			continue
		}
		if len(row)-2 == 0 {
			return Grid{}
		}
		newmp[i-1] = row[1 : len(row)-2]
	}
	var newhidden = make([][]bool, len(g.hidden)-2)
	for i, row := range g.hidden {
		if i == 0 || i >= len(newhidden) {
			continue
		}
		newhidden[i-1] = row[1 : len(row)-2]
	}
	return Grid{newmp, newhidden}
}
