package main

import (
	"aoc/lib"
	"fmt"
	"strings"
)

type Box struct {
	labels []string
	focals []int
}

func (b *Box) Remove(label string) {
	for i := range b.labels {
		if b.labels[i] == label {
			b.labels = append(b.labels[0:i], b.labels[i+1:]...)
			b.focals = append(b.focals[0:i], b.focals[i+1:]...)
			break
		}
	}
}

func (b *Box) Set(label string, focal int) {
	for i := range b.labels {
		if b.labels[i] == label {
			b.focals[i] = focal
			return
		}
	}
	b.labels = append(b.labels, label)
	b.focals = append(b.focals, focal)
}

func main() {
	line := lib.Lines()[0]
	instructions := strings.Split(line, ",")
	var part1 = 0
	var boxes = make([]Box, 256)
	for _, inst := range instructions {
		r := hash(inst)
		fmt.Println(inst, r)
		part1 += r
		label := inst[0:strings.IndexAny(inst, "=-")]
		box := hash(label)
		if inst[len(inst)-1] == '-' {
			boxes[box].Remove(label)
		} else {
			focal := lib.Int(inst[len(label)+1:])
			boxes[box].Set(label, focal)
		}
	}
	fmt.Println("part1", part1)
	fmt.Println("part2", lib.Sum(lib.MapIdx(boxes, sum)))

}

func hash(str string) int {
	val := 0
	for i := range str {
		val += int(byte(str[i]))
		val *= 17
		val %= 256
	}
	return val
}

func sum(b Box, idx int) (out int) {
	for i := range b.focals {
		out += (idx + 1) * (i + 1) * b.focals[i]
	}
	return
}
