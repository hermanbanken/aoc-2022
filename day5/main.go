package main

import (
	"aoc/lib"
	"bufio"
	"fmt"
	"sort"
	"strings"
)

type Range struct {
	Dst, Src, Len int
}

var conv [][]Range

func main() {
	r := lib.Reader()
	defer r.Close()
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	seeds := lib.Map(strings.Fields((scanner.Text()))[1:], lib.Int)
	conv = make([][]Range, 0)
	fmt.Println(seeds)
	var Map []Range
	for scanner.Scan() {
		t := scanner.Text()
		if strings.Contains(t, "-to-") {
			if len(Map) > 0 {
				conv = append(conv, Map)
			}
			Map = make([]Range, 0)
		} else if strings.TrimSpace(t) != "" {
			c := lib.Map(strings.Fields(t), lib.Int)
			Map = append(Map, Range{c[0], c[1], c[2]})
		}
	}
	conv = append(conv, Map)
	fmt.Println(conv)

	locations := []int{}
	for _, seed := range seeds {
		data := seed
		// fmt.Println("seed", data)
		for i := 0; i < len(conv); i++ {
			// fmt.Println(" ", data, lookup(data, i))
			data = lookup(data, i)
		}
		locations = append(locations, data)
	}
	fmt.Println(locations)
	sort.Ints(locations)
	fmt.Println(locations[0])
}

func lookup(nr int, mapIdx int) int {
	for _, r := range conv[mapIdx] {
		if nr >= r.Src && nr <= r.Src+r.Len {
			return nr + r.Dst - r.Src
		}
	}
	return nr
}
