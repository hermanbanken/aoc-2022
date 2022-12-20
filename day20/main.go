package main

import (
	"aoc/lib"
	"fmt"
)

type Num struct {
	originalIdx int
	value       int64
}

func main() {
	nums := []*Num{}
	idxs := []int{}
	idx := 0
	lib.EachLine(func(line string) {
		nums = append(nums, &Num{originalIdx: idx, value: lib.Int64(line) * 811589153})
		idxs = append(idxs, idx)
		idx++
	})

	circ := make([]*Num, len(nums))
	copy(circ, nums)

	print := func() {
		for i := range circ {
			fmt.Print(circ[i].value, ", ")
		}
		fmt.Println()
	}

	print()
	for round := 0; round < 10; round++ {
		for i := range nums {
			num := nums[i]
			pos := int64(indexOf(circ, num))
			// log.Println("pos", pos, "idxs[i]", idxs[i])
			newPos := (int64(pos) + num.value) % int64(len(nums)-1)
			if newPos < 0 {
				newPos += int64(len(nums) - 1)
			}
			if newPos == 0 {
				newPos = int64(len(nums) - 1)
			}
			dir := lib.Ternary(pos > newPos, int64(-1), int64(1))
			for p := pos; p != newPos; p += dir {
				if p < p+dir {
					swap(circ, idxs, int(p), int(p+dir))
				} else {
					swap(circ, idxs, int(p+dir), int(p))
				}
			}
			// print()
		}
	}
	print()

	for i := range circ {
		if circ[i].value == 0 {
			fmt.Println(circ[(i+1000)%len(circ)].value, circ[(i+2000)%len(circ)].value, circ[(i+3000)%len(circ)].value)
			fmt.Println(circ[(i+1000)%len(circ)].value + circ[(i+2000)%len(circ)].value + circ[(i+3000)%len(circ)].value)
		}
	}
	// not 1608 -8025 3733 => -2684
}

func swap(list []*Num, idxs []int, i, j int) {
	tmp := list[i]
	list[i] = list[j]
	list[j] = tmp
	idxs[i] += 1
	idxs[j] -= 1
}

func indexOf(list []*Num, el *Num) int {
	for pp, pointer := range list {
		if pointer == el {
			return pp
		}
	}
	panic("not found")
}
