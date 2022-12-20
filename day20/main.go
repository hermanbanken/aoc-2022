package main

import (
	"aoc/lib"
	"fmt"
)

type Num struct {
	originalIdx int
	value       int
}

func main() {
	nums := []*Num{}
	idxs := []int{}
	idx := 0
	lib.EachLine(func(line string) {
		nums = append(nums, &Num{originalIdx: idx, value: lib.Int(line)})
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
	for i := range nums {
		num := nums[i]
		pos := indexOf(circ, num)
		// log.Println("pos", pos, "idxs[i]", idxs[i])
		newPos := (pos + num.value) % (len(nums) - 1)
		if newPos < 0 {
			newPos += len(nums) - 1
		}
		if newPos == 0 {
			newPos = len(nums) - 1
		}
		dir := lib.Ternary(pos > newPos, -1, 1)
		for p := pos; p != newPos; p += dir {
			if p < p+dir {
				swap(circ, idxs, p, p+dir)
			} else {
				swap(circ, idxs, p+dir, p)
			}
		}
		// print()
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

func remove(list []*Num, idx int) []*Num {
	removed := append(list[0:idx], list[idx+1:]...)
	// log.Println(removed, "removed")
	return removed
}

func add(list []*Num, idx int, new *Num) []*Num {
	res := make([]*Num, len(list)+1)
	copy(res[0:idx], list[0:idx])
	res[idx] = new
	copy(res[idx+1:], list[idx:])
	// log.Println(res, "added")
	return res
}
