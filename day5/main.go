package main

import (
	"aoc/lib"
	"bufio"
	"fmt"
	"sort"
	"strings"
)

type Range struct {
	Src, Dst, Len int
}

func (r Range) String() string {
	return fmt.Sprintf("[%d-%d -> %d-%d]", r.Src, r.Src+r.Len, r.Dst, r.Dst+r.Len)
}

var conv [][]Range

func main() {
	r := lib.Reader()
	defer r.Close()
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	partTwo := true

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
			Map = append(Map, Range{Dst: c[0], Src: c[1], Len: c[2]})
		}
		sort.SliceStable(Map, func(i, j int) bool {
			if Map[i].Src == Map[j].Src {
				return Map[i].Len < Map[j].Len
			}
			return Map[i].Src+Map[i].Len < Map[j].Src+Map[j].Len
		})
	}
	conv = append(conv, Map)
	fmt.Println(conv)

	locations := []int{}
	for i := 0; i < len(seeds); i += lib.Ternary(partTwo, 2, 1) {
		start := seeds[i]
		var count int
		if partTwo {
			count = seeds[i+1]
		} else {
			count = 1
		}
		fmt.Println("seed", seeds[i], count)
		for j := start; j < start+count; j++ {
			data := j
			// fmt.Println("seed", data)
			for mapIdx := 0; mapIdx < len(conv); mapIdx++ {
				// fmt.Println(" ", data, lookup(data, mapIdx))
				data = lookup(data, mapIdx)
			}
			locations = append(locations, data)
		}
	}
	fmt.Println(locations)
	sort.Ints(locations)
	fmt.Println(locations[0])
}

func lookup2(nr int, mapIdx int) int {
	for _, r := range conv[mapIdx] {
		if nr >= r.Src && nr <= r.Src+r.Len {
			return nr + r.Dst - r.Src
		}
	}
	return nr
}

func lookup(nr int, mapIdx int) int {
	mp := conv[mapIdx]
	idx := sort.Search(len(mp), func(i int) bool {
		return mp[i].Src+mp[i].Len > nr
	})
	if idx < len(mp) {
		r := mp[idx]
		// fmt.Println("  ", mp[idx], nr, "=>", idx, "=>", nr+r.Dst-r.Src)
		if nr >= r.Src && nr < r.Src+r.Len {
			return nr + r.Dst - r.Src
		}
	} else {
		// fmt.Println("  not found", nr, idx, mp)
	}
	return nr
}
