package main

import (
	"aoc/lib"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//405
	lines := lib.Lines()
	start := 0
	end := 0
	sum := 0
	for ; end < len(lines); end++ {
		if lines[end] == "" {
			result := solve(lines[start:end])
			sum += result
			fmt.Println(result)
			start = end
		}
	}
	result := solve(lines[start+1:])
	sum += result
	fmt.Println(result)
	fmt.Println(sum)
}

func solve(lines []string) int {
	var hor = make([]int, len(lines))
	var ver = make([]int, len(lines[0]))
	var idx = make([]int, len(lines))
	var idx2 = make([]int, len(lines[0]))
	for i, l := range lines {
		hor[i], _ = strconv.Atoi(strings.Replace(strings.Replace(l, ".", "1", -1), "#", "0", -1))
		idx[i] = i
	}
	for j := range lines[0] {
		idx2[j] = j
		ver[j], _ = strconv.Atoi(strings.Replace(strings.Replace(line(lines, j), ".", "1", -1), "#", "0", -1))
	}
	sort.Slice(idx, func(i, j int) bool {
		return hor[i] > hor[j]
	})
	sort.Slice(idx2, func(i, j int) bool {
		return ver[i] > ver[j]
	})
	fmt.Println(idx, idx2)

	return 0
	if len(lines) == 0 {
		panic("invalid lines")
	}
	fmt.Println()
	for _, line := range lines {
		fmt.Println(line)
	}
	matches := []int{}
	for x := 1; x < len(lines[0]); x++ {
		if potentiallyMirrored(lines, x, 0) {
			matches = append(matches, x)
		}
	}
	for y := 1; y < len(lines); y++ {
		if potentiallyMirrored(lines, 0, y) {
			matches = append(matches, (y)*100)
		}
	}
	if len(matches) > 0 {
		sort.Sort(sort.Reverse(sort.IntSlice(matches)))
		fmt.Println("matches", matches)
		return matches[0]
	}
	panic("not mirrored")
}

func potentiallyMirrored(lines []string, x int, y int) bool {
	if x > 0 {
		x1 := x - 1
		x2 := x
		return line(lines, x1) == line(lines, x2)
	} else {
		y1 := y - 1
		y2 := y
		return lines[y1] == lines[y2]
	}
}

func areMirrored(lines []string, x int, y int) bool {
	if x > 0 {
		x1 := x - 1
		x2 := x
		for {
			if x1 < 0 || x2 > len(lines[0])-1 {
				return true
			}
			if line(lines, x1) == line(lines, x2) {
				x1 -= 1
				x2 += 1
			} else {
				return false
			}
		}
	} else {
		y1 := y - 1
		y2 := y
		for {
			if y1 < 0 || y2 > len(lines)-1 {
				return true
			}
			if lines[y1] == lines[y2] {
				y1 -= 1
				y2 += 1
			} else {
				return false
			}
		}
	}
}

func line(lines []string, x int) string {
	var chars []byte
	for _, line := range lines {
		chars = append(chars, line[x])
	}
	return string(chars)
}
